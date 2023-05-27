package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/DeveloperDurp/DurpCLI/cmd/auth"
	"gitlab.com/DeveloperDurp/DurpCLI/cmd/cfg"
	"gitlab.com/DeveloperDurp/DurpCLI/cmd/net"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "DurpCLI",
	Short: "CLI Tool made for Durp",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func setDefaults() {
	viper.SetDefault("auth.url", "https://authentik.durp.info/application/o/token/")
	viper.SetDefault("auth.grantType", "client_credentials")
	viper.SetDefault("auth.clientID", "")
	viper.SetDefault("auth.username", "")
	viper.SetDefault("auth.password", "")
}

func init() {
	cobra.OnInitialize(initConfig)

	setDefaults()

	err := viper.WriteConfigAs(".DurpCLI.yaml")
	if err != nil {
		fmt.Println(err)
	}

	rootCmd.AddCommand(net.NetCmd)
	rootCmd.AddCommand(auth.AuthCmd)
	rootCmd.AddCommand(cfg.Cfgcmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.DurpCLI.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".DurpCLI")
	}

	viper.AutomaticEnv()
	viper.ReadInConfig()

}
