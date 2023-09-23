package cmd

import (
	"github.com/Kavinraja-G/node-gizmo/pkg/cmd"
	"github.com/Kavinraja-G/node-gizmo/utils"
)

func init() {
	// inits the clientset and other generic configs if any
	utils.InitConfig()
}

// Execute drives the root 'nodegizmo' command
func Execute() error {
	root := cmd.NewCmdRoot()
	if err := root.Execute(); err != nil {
		return err
	}

	return nil
}
