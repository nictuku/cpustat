// Package cpustat provides access to CPU usage stats for the current process.
package cpustat

import (
	"sync"
	"syscall"
	"time"
)

type CPUStat struct {
	mu                sync.RWMutex
	prevCPUTimeNano   int64
	prevCPUCollection time.Time
}

func New() *CPUStat {
	cpu := new(CPUStat)
	cpu.ProcCPU() // Initalize the internal state.
	return cpu
}

// ProcCPU returns the number of CPU core/s used by the present process since
// ProcCPU was last called. It combines the time spent on kernel and user
// modes.
func (c *CPUStat) ProcCPU() (float64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	r := new(syscall.Rusage)
	if err := syscall.Getrusage(syscall.RUSAGE_SELF, r); err != nil {
		return 0, err
	}
	sum := r.Utime.Nano() + r.Stime.Nano()
	delta := sum - c.prevCPUTimeNano
	rate := float64(delta) / float64(now.Sub(c.prevCPUCollection).Nanoseconds())

	c.prevCPUTimeNano = sum
	c.prevCPUCollection = now

	return rate, nil
}
