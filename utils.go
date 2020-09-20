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
