package cgroups

type ThrottlingData struct {
   // Number of periods with throttling active
   Periods uint64 `json:"periods, omitempty"`
   // Number of periods when the container hit its throtling limit.
   ThrottledPeriods uint64 `json:"throttled_periods, omitempty"`
   // Aggregate time the container was throttled for in nanoseconds.
   ThrottledTime uint64 `json:"throttled_time,omitempty"`
}

// CPUUsage denotes the usage of the CPU
// All CPU stats are aggregate since container inception
type CpuUsage struct {
   // Total CPU time consumed
   // Units: nanoseconds
   TotalUsage uint64 `json:"total_usage, omitempty"`
   // Total CPU time consumed per core.
   // Units: nanoseconds
   PercpuUsage []uint64 `json:"percpu_usage,omitempty"`
   // CPU time consumed per core in kernel mode
   // Units: nanoseconds
   PercpuUsageInKernelmode []unit64 `json:"percpu_usage_in_kernelmode"`
   // CPU time consumed per core in user mode
   // Units: nanoseconds.
   PercpuUsageInUsermode []uint64 `json:"percpu_usage_in_usermode"`
   // Time spent by tasks of the cgroup in user mode.
   // Units: nanoseconds
   UsageInKernelmode uint64 `json:"usage_in_kernelmode"`
   // Time spent by tasks of the cgroup in user mode
   // Units: nanoseconds
   UsageInUsermode uint64 `json:"usage_in_usermode"`
}

type CpuStats struct {
   CpuUsage CpuUsage `json:"cpu_usage, omitempty"`
   ThrottlingData ThrottlingData `json:"throttling_data, omitempty"`
}

type MemoryData struct {
   Usage uint64 `json:"usage, omitempty"`
   MaxUsage uint64 `json:"max_usage, omitempty"`
   Failcnt unit64 `json:"failcnt"`
   Limit uint64 `json:"limit"`
}

type MemoryStats struct {
   // memory used for cache
   Cache uint64 `json:"cache, omitempty"`
   // Usage of Memory
   Usage MemoryData `json:"usage, omitempty"`
   // Usage of memory + swap
   SwapUsage MemoryData  `json:"swap_usage, omitempty"`
   // Usage of kernel tcp memory
   kernelUsage MemoryData `json:"kernel_usage, omitempty"`
   // Usage of kernel tcp memory
   kernelTCPUsage MemoryData `json:"kernel_usage, omitempty"`
   // usage of memory pages by NUMA mode
   // see chapter 5.6 of memory controller discussion
   pageUsageByNUMA pageUsaeByNuma `json: "page usage by numa, omitempty"`
   // if true, memory usage is accounted for throuout a hireachy of cgroups
   UseHirearchy bool `json:"use_hirearcy"
   tats map[string]uint64 `json:"stats, omitempty"`
}
useHirearchy bool `json:"use_hirearchy"`o
