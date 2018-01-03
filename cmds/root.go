package cmds

import (
	"flag"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "ssl [command]",
		Short:             `ssl practice`,
		DisableAutoGenTag: true,
	}
	// ref: https://github.com/kubernetes/kubernetes/issues/17162#issuecomment-225596212
	flag.CommandLine.Parse([]string{})

	cmd.AddCommand(NewCmdServer())
	cmd.AddCommand(NewCmdClient())
	return cmd
}
