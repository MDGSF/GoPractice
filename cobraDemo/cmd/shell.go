package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var c2ID string

func init() {
	shellCmd.Flags().StringVar(&c2ID, "t", "", "c2 id")
}

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "shell devices",
	Long:  "long",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shellCmd c1 run args = ", args)
		fmt.Println("shellCmd cfgFile = ", cfgFile)
		fmt.Println("shellCmd c2ID = ", c2ID)
	},
}
