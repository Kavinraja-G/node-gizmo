package cmd

import (
	"os"

	"github.com/spf13/pflag"

	"github.com/Kavinraja-G/node-gizmo/pkg/cmd"
	"k8s.io/cli-runtime/pkg/genericiooptions"
)

func Execute() error {
	flags := pflag.NewFlagSet("kubectl-ng", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := cmd.NewCmdRoot(genericiooptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		return err
	}

	return nil
}
