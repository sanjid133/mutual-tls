package cmds

import (
	"github.com/spf13/cobra"
	"github.com/sanjid133/mutual-tls/pkg/server"
	"github.com/appscode/go/term"
)

func NewCmdServer() *cobra.Command  {
	cmd := &cobra.Command{
		Use: "server",
		Short: "run the server",
		DisableAutoGenTag:true,
		Run: func(cmd *cobra.Command, args []string) {
				if err := server.Init(); err != nil {
					term.Fatalln(err)
				}

		},
	}
	return cmd
}