package user

import "os"

func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown_hostname"
	}

	return hostname
}
