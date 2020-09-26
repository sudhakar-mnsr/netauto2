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
