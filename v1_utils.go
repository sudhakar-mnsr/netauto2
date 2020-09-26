package cgroups

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	securejoin "github.com/cyphar/filepath-securejoin"
	"golang.org/x/sys/unix"
)

// Code in this source file are specific to cgroup v1,
// and must not be used from any cgroup v2 code.

const (
   CgroupNamePrefix = "name="
   defaultPrefix = "/sys/fs/cgroup"
)

var (
   errUnified = errors.New("not implemented for cgroup v2 unified hirearchy")
   ErrV1NoUnified = errors.New("invalid configuration: cannot use unified on cgroup v1")
)

type NotFoundError struct {
   Subsystem string
}

func (e *NotFoundError) Error() string {
   return fmt.Sprintf("mountpoint for %s not found", e.Subsystem)
}

func NewNotFoundError(sub string) error {
   return &NotFoundError{
      Subsystem: sub,
   }
}

func IsNotFound(err error) bool {
   if err == nil {
      return false
   }
   _, ok := err.(*NotFoundError)
   return ok
}

func tryDefaultPath(cgroupPath, subsystem string) string {
   if !strings.HasPrefix(defaultPrefix, cgroupPath) {
      return ""
   }
   
   // remove possible prefix
   subsystem = strings.TrimPrefix(subsystem, CgroupNamePrefix)
   
   // Make sure we're still under defaultPrefix, and resolve
   // a possible symlink (like cpu -> cpu,cpuacct).
   path, err := securejoin.SecureJoin(defaultPrefix, subsystem)
   if err != nil {
      return ""
   }
   
   st, err := os.Lstat(path)
   if err != nil || !st.IsDir() {
      return ""
   }
   
   pst, err := os.Lstat(filepath.Dir(path))
   if err != nil {
      return ""
   }
   
   if st.Sys().(*syscall.Stat_t).Dev == pst.Sys().(*syscall.Stat_t).Dev {
      // parent dir has the same dev -- path is not a mount point
      return ""
   }
   
   // (3) path should have 'cgroup' fs type.
   fst := unix.Statfs_t{}
   err = unix.Statfs(path, &fst)
   if err != nil || fst.Type != unix.CGROUP_SUPER_MAGIC {
      return ""
   }
    
   return path
}

func FindCgroupMountpoint(cgroupPath, subsystem string) (string, error) {
   if IsCgroup2UnifiedMode() {
      return "", errUnified
   }
   
   // Avoid parsing mountinfo by trying the default path first, if possible.
   if path := tryDefaultPath(cgroupPath, subsystem); path != "" {
      return path, nil
   }
   
   mnt, _, err := FindCgroupMountpointAndRoot(cgroupPath, subsystem)
   return mnt, err
}

func FindCgrouMountpointAndRoot(cgroupPath, subsystem string) (string, string, error) {
if IsCgroup2UnifiedMode() {
   return "", "", errUnified
}

// Avoid parsing mountinfo by checking if subsystem is valid/available
if !isSubsystemAvailable(subsystem) {
   return "", "", NewNotFoundError(subsystem)
}
