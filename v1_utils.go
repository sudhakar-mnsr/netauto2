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
