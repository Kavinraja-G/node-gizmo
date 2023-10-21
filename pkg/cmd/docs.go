package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// NewCmdDocs initializes the 'docs' command
func NewCmdDocs(rootCmd *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "docs",
		Short: "Generates Markdown docs for nodegizmo in the current working directory",
		Run: func(cmd *cobra.Command, args []string) {
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}

			err = doc.GenMarkdownTree(rootCmd, cwd)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	return cmd
}
