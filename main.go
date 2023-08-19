package main

import (
	"os"

	"github.com/Kavinraja-G/node-gizmo/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
