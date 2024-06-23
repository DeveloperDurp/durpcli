package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cli/oauth/device"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

var generateTokenCmd = &cobra.Command{
	Use:   "generateToken",
	Short: "Prints disk usage of the current directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if clientID == "" {
			clientID = viper.GetViper().GetString("auth.clientID")
		}
		generateToken(clientID)
	},
}

func init() {
	generateTokenCmd.Flags().StringVarP(&clientID, "clientID", "c", "", "The ClientID")
	generateTokenCmd.Flags().
		StringVarP(&grantType, "grantType", "g", "client_credentials", "The Grant Type")
	generateTokenCmd.Flags().StringVarP(&url, "url", "", "", "Token URL")

	AuthCmd.AddCommand(generateTokenCmd)
}

func generateToken(
	clientID string,
) {
	scopes := []string{"openid", "email", "profile"}
	httpClient := http.DefaultClient

	code, err := device.RequestCode(httpClient, "https://authentik.durp.info/application/o/device/", clientID, scopes)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Copy code: %s\n", code.UserCode)
	fmt.Printf("then open: %s\n", code.VerificationURIComplete)

	accessToken, err := device.Wait(context.TODO(), httpClient, "https://authentik.durp.info/application/o/token/", device.WaitOptions{
		ClientID:   clientID,
		DeviceCode: code,
		GrantType:  "urn:ietf:params:oauth:grant-type:device_code",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Access token: %s\n", accessToken.Token)
}
