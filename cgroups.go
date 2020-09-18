package cgroups

import (
   "github.com/opencontainers/runc/libcontainer/configs"
)

type Manager interface {
// Applies cgroup configuration to the process with the specified pid
Apply(pid int) error

// Returns the PID's inside the cgroup set
GetPids() ([]int, error)

// Returns the PID's inside the cgroup set & all sub-cgroups
GetAllPids() ([]int, error)
