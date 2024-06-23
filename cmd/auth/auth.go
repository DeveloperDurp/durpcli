package auth

import (
	"github.com/spf13/cobra"
)

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	IDToken     string `json:"id_token"`
}

var (
	clientID  string
	grantType string
	url       string
	username  string
	password  string
	scope     string
)

var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "All things Authorization",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
