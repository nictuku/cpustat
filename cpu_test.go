package cpustat

import (
	"testing"
)

func TestProcCPU(t *testing.T) {
	cpu := New()
	for i := 0; i < 5; i++ {
		c, err := cpu.ProcCPU()
		if err != nil {
			t.Fatalf("ProcCPU failed: %v", err)
		}
		t.Logf("CPU: %.1f (zero is normal)", c)
	}
}
