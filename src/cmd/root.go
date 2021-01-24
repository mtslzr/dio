package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "dio",
	Short: "Command-line tool to bootstrap new projects.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Dio!")
	},
}

// Executes runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}