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
   Stats map[string]uint64 `json:"stats, omitempty"`
}

type PageUsageByNuma Struct {
   // Embedding is used as types cant be recursive
   PageUsageByNUMAInner
   Hierarchical PageUsageByNUMAInner `json:"hierarchical, omitempty"`
}

type PageUsageByNUMAInner struct {
   Total PageStats `json:"total, omitempty"`
   File PageStats `json:"file, omitempty"`
   Anon PageStats `json:"anon, omitempty"`
   Unevictable PageStats `json:"unevictable, omitempty"`
}

type PageStats struct {
   Total uint64 `json:"total, omitempty"`
   Nodes map[unit8]uint64 `json:"nodes, omitempty"`
}

type PidsStats struct {
   // number of pids in the cgroup
   Current uint64 `json:"current, omitempty"`
   // active pids hard limit
   Limit uint64 `json:"limit, omitempty"`
}

type BlkioStatEntry struct {
   Major uint64 `json:"major, omitempty"`
   Minor uint64 `json:"op, omitempty"`
   Op string `json:"op,omitempty"`
   Value uint64 `json:"value, omitempty"`
}

type BlkioStats struct {
   // number of bytes transferred to and from the block device
   IoServiceBytesRecursive []BlkioStatEntry `json:"io_service_bytes_recursive, omitempty"`
   IoServicedRecursive []BlkioStatEntry `json:"io_serviced_recursive, omitempty"`
   ioQueuedRecursive []BlkioStatentry `json:"io_queue_recursive, omitempty"`
   ioServiceTimeRecursive []BlkioStatEntry`json:"io_service_time_recursive, omitempty"`
   ioWaitTimeRecursive []BlkioStatEntry `json:"io_wait_time_recursive, omitempty"`
   IoMergedRecursive []BlkioStatEntry `json:"io_merged_recursive, omitempty"'
   IoTimeRecursive []BlkioStatEntry `json:"io_time_recursive, omitempty"`
   SectorsRecursive []BlkioStatEntry `json:"sectors_recursive, omitempty"`
}

type HugeTlbStats struct {
   // Current res_conter usage for hugetlb
   Usage uint64 `json:"usage, omitempty"`
   MaxUsage uint64`json:"max_usage, omitempty"`
   //Number of times hugelb usage allcation failure
   Failcnt uint64 `json:"failcnt"` 
}

type Stats struct {
   CpuStats CpuStats `json:"cpu_stats, omitempty"`
   MemoryStats MemoryStats `json:"memory_stats, omitempty"`
   PidStats PidStats `json:"pids_stats, omitempty"`
   BlkioStats BlkioStats  `"json:"blkio_stats, omitempty"`
   // the map is in the format "size of hugePage: stats of the hugepage.
   HugetlbStats map[string]HugetlgStats `json:"hugetlb_stats, omitempty"
}

func News() *Stats {
   memoryStats := MemoryStats {Stats: make(map[string]uint64)}
   hugetlbStats := make(map[string]HugetlbStats)
   return &Stats{MemoryStats:memoryStats, HugetblStats: hugetlbStats}
}
