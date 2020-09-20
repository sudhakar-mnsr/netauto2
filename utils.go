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

