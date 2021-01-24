package cmd

import (
	"github.com/mtslzr/dio/src/pkg/auth"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

const (
	flagTokenDesc   = "personal access token for Github"
	flagUserDesc    = "username for Github"
)

type Flags struct {
	Destroy bool   `json:"destroy"`
	Status  bool   `json:"status"`
	Token   string `json:"token"`
	User    string `json:"user"`
}

var (
	flags = Flags{}

	authCmd = &cobra.Command{
		Use:   "auth",
		Short: "Authenticate Dio with Github.",
		Run: func(cmd *cobra.Command, args []string) {
			err := auth.Authenticate(flags.Token, flags.User)
			if err != nil {
				returnResult(err)
			}

			returnResult(auth.Status())
		},
	}

	authDelCmd = &cobra.Command{
		Use: "delete",
		Short: "Delete existing token configuration.",
		Run: func(cmd *cobra.Command, arg []string) {
			returnResult(auth.Destroy())
		},
	}

	authStatCmd = &cobra.Command{
		Use: "status",
		Short: "Check existing token configuration.",
		Run: func(cmd *cobra.Command, arg []string) {
			returnResult(auth.Status())
		},
	}
)

func init() {
	authCmd.Flags().StringVarP(&flags.Token, "token", "t", "", flagTokenDesc)
	authCmd.Flags().StringVarP(&flags.User, "user", "u", "", flagUserDesc)
	authCmd.MarkFlagRequired("token")
	authCmd.MarkFlagRequired("user")

	authCmd.AddCommand(authDelCmd)
	authCmd.AddCommand(authStatCmd)
	rootCmd.AddCommand(authCmd)
}

func returnResult(err error) {
	if err != nil {
		log.Errorf("Encountered errors while processing request.")
		os.Exit(1)
	}
	os.Exit(0)
}
