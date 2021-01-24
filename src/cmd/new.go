package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use: "new",
	Short: "Create new project using Dio.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("New project with Dio!")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}