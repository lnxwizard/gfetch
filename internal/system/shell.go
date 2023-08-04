package system

import (
	"os"
	"path"
)

func GetShell() string {
	shell := "sh"
	shellPath := "/usr/bin/sh"

	shellPath, status := os.LookupEnv("SHELL")
	if status {
		shell = path.Base(shellPath)
	}

	return shell
}
