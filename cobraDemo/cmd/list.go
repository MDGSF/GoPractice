package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list devices",
	Long:  "long",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listCmd c1 run args = ", args)
		fmt.Println("listCmd cfgFile = ", cfgFile)
	},
}
