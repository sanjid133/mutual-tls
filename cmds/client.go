package cmds


import (
	"github.com/spf13/cobra"
	"github.com/sanjid133/mutual-tls/pkg/client"
	"fmt"
)

func NewCmdClient() *cobra.Command  {
	cmd := &cobra.Command{
		Use: "client",
		Short: "run the server",
		DisableAutoGenTag:true,
		Run: func(cmd *cobra.Command, args []string) {
			if err := client.Run(); err != nil {
				fmt.Println(err)
			}

		},
	}
	return cmd
}
