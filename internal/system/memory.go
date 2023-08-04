package system

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetMemoryInfo() string {
	v, err := mem.VirtualMemory()
	if err != nil {
		return "unknown_memory"
	}

	return fmt.Sprintf("%d / %d MiB", (v.Used / 1048576), (v.Total / 1048576))
}
