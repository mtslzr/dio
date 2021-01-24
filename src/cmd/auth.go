package cmd

import (
	"github.com/mtslzr/dio/src/pkg/auth"
	"github.com/spf13/cobra"
)

const (
	flagDestroyDesc = "destroy any existing token settings"
	flagStatusDesc  = "check existing token settings"
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
			if flags.Status {
				auth.Status()
				return
			} else if flags.Destroy {
				auth.Destroy()
				auth.Status()
				return
			}

			if flags.Token != "" {
				auth.Authenticate(flags.Token)
			}

			auth.Status()
			return
		},
	}
)

func init() {
	authCmd.Flags().BoolVarP(&flags.Destroy, "destroy", "d", false, flagDestroyDesc)
	authCmd.Flags().BoolVarP(&flags.Status, "status", "s", false, flagStatusDesc)
	authCmd.Flags().StringVarP(&flags.Token, "token", "t", "", flagTokenDesc)
	authCmd.Flags().StringVarP(&flags.User, "user", "u", "", flagUserDesc)

	rootCmd.AddCommand(authCmd)
}
