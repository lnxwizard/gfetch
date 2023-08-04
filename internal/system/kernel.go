package system

import "github.com/shirou/gopsutil/v3/host"

func GetKernel() string {
	h, err := host.Info()
	if err != nil {
		return "unknown_kernel"
	}

	return h.KernelVersion
}
