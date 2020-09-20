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
