package cmd

import (
	"github.com/Kavinraja-G/node-gizmo/pkg/cmd"
)

// Execute drives the root 'nodegizmo' command
func Execute() error {
	root := cmd.NewCmdRoot()
	if err := root.Execute(); err != nil {
		return err
	}

	return nil
}
