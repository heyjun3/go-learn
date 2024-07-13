package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initVersionCmd())
}

func initVersionCmd() *cobra.Command {
	var version int

	var versionCmd = &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("version %d\n", version)
			fmt.Println("cmd", cmd)
			fmt.Println("args", args)
		},
	}

	versionCmd.Flags().IntVarP(&version, "version", "v", 1, "expect version")
	return versionCmd
}
