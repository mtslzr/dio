package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "dio",
	Short: "Command-line tool to bootstrap new projects.",
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

// Executes runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}