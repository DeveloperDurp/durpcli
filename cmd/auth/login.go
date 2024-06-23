package auth

import (
	"context"
	"log"
	"net/http"
	"os/exec"
	"os/user"

	"github.com/cli/oauth/device"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zalando/go-keyring"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if clientID == "" {
			clientID = viper.GetViper().GetString("auth.clientID")
		}
		login(clientID)
	},
}

func init() {

	AuthCmd.Flags().StringVarP(&clientID, "clientID", "c", "", "The ClientID")
	AuthCmd.AddCommand(loginCmd)
}

func login(clientID string) {

	service := "durpcli"
	user, _ := user.Current()

	scopes := []string{"openid", "email", "profile"}
	httpClient := http.DefaultClient

	code, err := device.RequestCode(httpClient, "https://authentik.durp.info/application/o/device/", clientID, scopes)
	if err != nil {
		panic(err)
	}

	exec.Command("open", code.VerificationURIComplete).Run()

	accessToken, err := device.Wait(context.TODO(), httpClient, "https://authentik.durp.info/application/o/token/", device.WaitOptions{
		ClientID:   clientID,
		DeviceCode: code,
		GrantType:  "urn:ietf:params:oauth:grant-type:device_code",
	})
	if err != nil {
		panic(err)
	}

	err = keyring.Set(service, user.Username, accessToken.Token)
	if err != nil {
		log.Fatal(err)
	}

}
