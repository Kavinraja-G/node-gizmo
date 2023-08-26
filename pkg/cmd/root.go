package cmd

import (
	"github.com/Kavinraja-G/node-gizmo/pkg/cmd/nodes"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericiooptions"
)

func NewCmdRoot(streams genericiooptions.IOStreams) *cobra.Command {
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

	return cmd
}
