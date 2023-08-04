package main

import (
	rootCmd "gfetch/pkg/cmd/root"
	"os"
)

func main() {
	if err := rootCmd.NewCmdRoot().Execute(); err != nil {
		os.Exit(1)
	}
}
