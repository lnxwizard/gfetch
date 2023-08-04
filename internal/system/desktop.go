package system

import "os"

func GetDE() string {
	if env, status := os.LookupEnv("DESKTOP_SESSION"); status {
		return env
	} else if env, status = os.LookupEnv("XDG_SESSION_DESKTOP"); status {
		return env
	}

	return "unknown_desktop_environment"
}
