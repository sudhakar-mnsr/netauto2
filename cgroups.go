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

// Returns statistics for the cgroup set
GetStats() (*Stats, error)

// Toggles the freezer cgroup according with specified state
Freeze(state configs.FreezerState) error

// Destroys the cgroup set
Destroy() error

// Path returns a cgroup path to the specified controller/subsystem
// For cgroupv2, the argument is unused and can be empty.
Path(string) string

// Sets the cgroup as configured.
Set(container *configs.Config) error

