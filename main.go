package main

import (
	"fmt"
	"os"

	"./cmd"
)

var (
	version = "dev"
	commit  = "unknown"
)

func main() {
	cmd.ConfigureVersion(version, commit)

	err := cmd.RootCmd.Execute()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
