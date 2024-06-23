package auth

import (
	"os/user"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logout()
	},
}

func init() {

	AuthCmd.AddCommand(logoutCmd)
}

func logout() {

	service := "durpcli"
	user, _ := user.Current()

	keyring.Delete(service, user.Username)
}
