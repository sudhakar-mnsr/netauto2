package cgroups

type ThrottlingData struct {
   // Number of periods with throttling active
   Periods uint64 `json:"periods, omitempty"`
   // Number of periods when the container hit its throtling limit.
   ThrottledPeriods uint64 `json:"throttled_periods, omitempty"`
   // Aggregate time the container was throttled for in nanoseconds.
   ThrottledTime uint64 `json:"throttled_time,omitempty"`
}
