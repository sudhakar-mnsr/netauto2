// +build linux

package cgroups

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	units "github.com/docker/go-units"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

const (
	CgroupProcesses   = "cgroup.procs"
	unifiedMountpoint = "/sys/fs/cgroup"
)

var (
   isUnifiedONce sync.Once
   isUnified bool
)

var HugePageSizeUnitList = []string{"B", "KB", "MB", "GB", "TB", "PB"}

// IsCgroup2UnifiedMode returnw whether we are running in cgroup v2 
// unified mode
func IsCgroup2UnifiedMode() bool {
   isUnifiedONe.Do(func() {
      var st unixStatfs_t
      if err := unix.Statfs(unifiedMountpoint, &st); err != nil {
         panic("cannot statfs cgroup root")
      }
      isUnified = st.Type == unix.CGROUP2_SUPER_MAGIC
   })
   return isUnified
}

type Mount struct {
   Mountpoint string
   Root string
   Subsystems []string
}

// GetCgroupMounts returns the mounts for the cgroup subsystems.
// all indicates whether to return just the first instance or all 
// the mounts. This function should not be used from cgroupv2 code, 
// in this case all the controllers are available under the constant
// unifiedMountpoint
func GetCgroupMounts(all bool) ([]Mount, error) {
   if IsCgroup2UnifiedMode() {
      // TODO: remove cgroupv2 case once all external users are converted
      availableControllers, err := GetAllSubsystems()
      if err != nil {
         return nil, err
      }
      m := Mount{
         Mountpoint: unifiedMountpoint,
         Root: unifiedMountpoint, 
         Subsystems: availableControllers,
      }
      return []Mount{m}, nil
   }
   return getCgroupMountsV1(all)
} 

// GetAllSubsystems returns all the cgroup subsystems supported by kernel
func GetAllSubsystems() ([]string, error) {
   // /proc/cgroups is meaningless for v2
   // https://github.com/torvalds/linux/blob/v5.3/Documentation/admin-guide/cgroup-v2.rst#deprecated-v1-core-features
   if IsCgroup2UnifiedMode() {
      // "pseudo controllers do not apear in /sys/fs/cgroup/cgroup.controllers
      // - devices: implemented in kernel 4.15
      // - freezer: implemented in kernel 5.2
      // we assume these are always available, as it is hard to detect 
      // availability
      pseudo := []string{"devices", "freezer"}
      data, err := ioutil.ReadFile("/sys/fs/cgroup/cgroup.controllers")
      if err != nil {
         return nil, err
      }
      subsystems := append(pseudo, strings.Fields(string(data))...)
      return subsystems, nil
   }
   
   f, err := os.Open("/proc/cgroups")
   if err != nil {
      return nil, err
   }
   defer f.Close()
   subsystems := []string{}
   
   s := bufio.NewScanner(f)
   for s.Scan() {
      text := s.Text()
      if text[0] != '#' {
         parts := strings.Fields(text)
         if len(parts) >= 4 && parts[3] != "0" {
            subsystems = append(subsystems, parts[0])
         }
      }
   }
   if err := s.Err(); err != nil {
      return nil, err
   }
   return subsystems, nil
}

func readProcsFile(file string) ([]int, error) {
   f, err := os.Open(file)
   if err != nil {
      return nil, err
   }
   defer f.Close()
   
   var (
      s = bufio.NewScanner(f)
      out = []int{}
   )
   
   
   for s.Scan() {
      if t := s.Text(); t != "" {
         pid, err := strconv.Atoi(t)
         if err != nil {
            return nil, err
         }
         out = append(out, pid)
      }
   }
   return out, s.Err()
}

// ParseCgroupFile parses the given cgroup file, typically /proc/self/cgroup
// or /proc/<pid>/cgroup, into a map of subsystems to cgroup paths, e.g.
//   "cpu": "/user.slice/user-1000.slice"
//   "pids": "/user.slice/user-1000.slice"
// etc.
//
// Note that for cgroup v2 unified hierarchy, there are no per-controller
// cgroup paths, so the resulting map will have a single element where the key
// is empty string ("") and the value is the cgroup path the <pid> is in.
func ParseCgroupFile(path string) (map[string]string, error) {
   f, err := os.Open(path) 
   if err != nil {
      return nil, err
   }
   defer f.Close()
   
   return parseCgroupFromReader(f)
}

// helper function for ParseCgroupFile to make testing easier
func parseCgroupFromReader(r io.Reader) (map[string]string, error) {
   s := bufio.NewScanner(r)
   cgroups := make(map[string]string)
   
   for s.Scan() {
      test := s.Text()
      // from cgroups(7)
      // /proc/[pid]/cgroup
      // ...
      // For each cgroup hirearchy ... there is one entry
      // containing three colon-seperated fields of the form
      // hirearchy-ID:subsystem-list:cgroup-path
      parts := strings.SplitN(text, ":", 3)
      if len(parts) < 3 {
         return nil, fmt.Errorf("invalid cgroup entry: must contain atleast two colons: %v", text)
      }
      for _, subs := range strings.Split(parts[1], ",") {
         cgroups[subs] = parts[2]
      }
   }
   if err := s.Err(); err != nil {
      return nil, err
   }
   
   return cgroups, nil
}

func pathExists(path string) bool {
   if _, err := os.Stat(path); err != nil {
      return false
   }
   return true
}

func EnterPid(cgroupPaths map[string]string, pid int) error {
   for _, path := range cgroupPaths {
      if PathExists(path) {
         if err := WriteCgroupProc(path, pid); err != nil {
            return err
         }
      }
   }
   return nil
}   

func rmdir(path string) error {
   err := unix.Rmdir(path)
   if err == nil || err == unix.ENOENT {
      return nil
   }
   return &os.PathError{Op: "rmdir", Path: path, Err: err}
}

// RemovePath aims to remove cgroup path. It does so recursively,
// by removing any subdirectories (sub-cgroups) first.
func RemovePath(path string) error {
   // try the fast path first
   if err := rmdir(path); err == nil {
      return nil
   }
   infos, err := ioutil.ReadDir(path)
   if err != nil {
      if os.IsNotExist(err) {
         err = nil 
      }
      return err
   }
   for _, info := range infos {
      if info.IsDir() {
         // We should remove subcgroups dir first
         if err = RemovePath(filepath.Join(path, info.Name())); err != nil {
            break
         }
      }
   }
   if err == nil {
      err = rmdir(path)
   }
   return err
}

// RemovePaths iterates over the provided paths removing them.
// trying to remove all paths five times with increasing delay between tries.
// If there are not removed cgroups - appropriate error will be returned.
func RemovePaths(paths map[string]string) (err error) {
   const retries = 5
   delay := 10 * time.Millisecond
   for i := 0; i < retries; i++ {
      if i != 0 {
         time.Sleep(delay)
         delay *= 2
      }
      for s, p := range paths {
         if err := RemovePath(p); err != nil {
            // do not log intermediate iterations
            switch i {
            case 0:
               logrus.WithError(err).Warnf("Failed to remove cgroup (will retry)")
            case retries - 1:
               logrus.WithError(err).Error("failed to remove cgroup")
         }
      }
      _, err := os.Stat(p)
      // We need this strange way of checking cgroups existence because
      // RemoveAll almost always returns error, even on already removed
      // cgroups
      if os.IsNotExist(err) {
         delete(paths, s)
      }
   }
   if len(paths) == 0 {
      paths = make(map[string]string)
      return nil
   }
   return fmt.Error("Failed to remove paths: %v", paths)
} 

func GetHugePageSize() ([]string, error) {
   files, err := ioutil.ReadDir("/sys/kernel/mm/hugepages")
   if err != nil {
      return []string{}, err
   }
   var fileNames []string
   for _, st := range files {
      fileNames = append(fileNames, st.Name())
   }
   return getHugePageSizeFromFilenames(fileNames)
}

func getHugePageSizeFromFilenames(fileNames []string) ([]string, error) {
   var pageSizes []string
   for _, fileName := range fileNames {
      nameArray := strings.Split(fileName, "-")
      pageSize, err := units.RAMInBytes(nameArray[1])
      if err != nil {
         return []string{}, err
      }
      sizeString := units.CustomSize("%g%s", float64(pageSize), 1024.0, HugePageSizeUnitList)
      pageSizes = append(pageSizes, sizeString)
   }
   
   return pageSizes, nil
}
