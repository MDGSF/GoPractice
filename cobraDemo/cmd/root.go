package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

func init() {
	//cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(shellCmd)
}

var rootCmd = &cobra.Command{
	Use:   "c1",
	Short: "c1 is a test",
	Long:  "c1 is a long test",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rootCmd c1 run args = ", args)
		fmt.Println("rootCmd cfgFile = ", cfgFile)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
