package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

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
		if grantType == "" {
			grantType = viper.GetViper().GetString("auth.grantType")
		}
		if url == "" {
			url = viper.GetViper().GetString("auth.url")
		}
		if username == "" {
			username = viper.GetViper().GetString("auth.username")
		}
		if password == "" {
			password = viper.GetViper().GetString("auth.password")
		}

		generateToken(clientID, grantType, url, username, password)
	},
}

func init() {
	generateTokenCmd.Flags().StringVarP(&clientID, "clientID", "c", "", "The ClientID")
	generateTokenCmd.Flags().
		StringVarP(&grantType, "grantType", "g", "client_credentials", "The Grant Type")
	generateTokenCmd.Flags().StringVarP(&url, "url", "", "", "Token URL")
	generateTokenCmd.Flags().StringVarP(&username, "username", "u", "", "username")
	generateTokenCmd.Flags().StringVarP(&password, "password", "p", "", "password")
	generateTokenCmd.Flags().StringVarP(&scope, "scope", "s", "openid profile", "scope")

	AuthCmd.AddCommand(generateTokenCmd)
}

func generateToken(
	clientID string,
	grantType string,
	url string,
	username string,
	password string,
) {
	formData := fmt.Sprintf("grant_type=%s&client_id=%s&username=%s&password=%s&scope=%s",
		grantType, clientID, username, password, scope)

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(formData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var response accessTokenResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	fmt.Println("Access_Token:", response.AccessToken)
	fmt.Println("Token_Type:", response.TokenType)
	fmt.Println("Expires_In:", response.ExpiresIn)
	fmt.Println("ID_Token:", response.IDToken)
}
