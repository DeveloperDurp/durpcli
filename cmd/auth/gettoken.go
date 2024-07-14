package auth

import (
	"fmt"
	"log"
	"os/user"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var getTokenCmd = &cobra.Command{
	Use:   "get token",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		token := gettoken()

		fmt.Println(token)
	},
}

func init() {

	AuthCmd.AddCommand(getTokenCmd)
}

func gettoken() string {

	service := "durpcli"
	user, _ := user.Current()

	token, err := keyring.Get(service, user.Username)
	if err != nil {
		log.Fatal(err)
	}

	return token
}
