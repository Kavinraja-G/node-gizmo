package cmd

import (
	"github.com/Kavinraja-G/node-gizmo/pkg/cmd/nodepool"
	"github.com/Kavinraja-G/node-gizmo/pkg/cmd/nodes"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"log"
)

// NewCmdRoot initializes the root command 'nodegizmo'
func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nodegizmo",
		Aliases: []string{"ng"},
		Short:   "Nodegizmo - A CLI utility for your Kubernetes nodes",
		RunE: func(c *cobra.Command, args []string) error {
			if err := c.Help(); err != nil {
				return err
			}
			return nil
		},
	}

	// child commands
	cmd.AddCommand(nodes.NewCmdNodeInfo())
	cmd.AddCommand(nodepool.NewCmdNodepoolInfo())

	// generate markdown docs
	err := doc.GenMarkdownTree(cmd, "docs")
	if err != nil {
		log.Fatalf("Error generating md docs for nodegizmo: %v", err)
	}

	return cmd
}
