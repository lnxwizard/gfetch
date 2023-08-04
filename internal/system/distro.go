package system

import (
	"os"
	"strings"
)

func GetDistro() string {
	contents, err := os.ReadFile("/etc/lsb-release")
	if err != nil {
		return "unknown_distro"
	}

	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Split(line, "=")
		if len(fields) == 2 && fields[0] == "DISTRIB_DESCRIPTION" {
			return strings.Trim(fields[1], `"`)
		}
	}

	return "unknown_distro"
}
