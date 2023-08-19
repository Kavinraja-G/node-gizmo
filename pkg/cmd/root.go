package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericiooptions"
)

func NewCmdRoot(streams genericiooptions.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nodegizmo node [flags]",
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
	cmd.AddCommand(NewCmdNodeInfo())

	return cmd
}
