package cmd

import (
	"fmt"
	"os"
	"os/user"

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

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}

func setDefaults() {
	viper.SetDefault("auth.url", "https://authentik.durp.info/application/o/token/")
	viper.SetDefault("auth.grantType", "client_credentials")
	viper.SetDefault("auth.clientID", "")
	viper.SetDefault("auth.username", "")
	viper.SetDefault("auth.password", "")
}

func init() {
	//if cfgFile != "" {
	//	viper.SetConfigFile(cfgFile)
	//} else {
	//	cobra.CheckErr(err)
	//}
	setDefaults()
	initConfig()
	loadConfig()

	//	styles := tui.DefaultStyles()
	//	styles.Title.BorderForeground(lipgloss.AdaptiveColor{Light: `#E3BD2D`, Dark: `#E3BD2D`})
	//	styles.Border.BorderForeground(lipgloss.AdaptiveColor{Light: `#E3BD2D`, Dark: `#E3BD2D`})
	//	styles.SelectedItem.Foreground(lipgloss.AdaptiveColor{Light: `#353C3B`, Dark: `#353C3B`}).
	//		Background(lipgloss.AdaptiveColor{Light: `#E3BD2D`, Dark: `#E3BD2D`})
	//
	//	b := tui.New(tui.WithStyles(styles))

	//	rootCmd.SetUsageFunc(b.UsageFunc)
	//	rootCmd.SetHelpFunc(b.HelpFunc)

	rootCmd.AddCommand(net.NetCmd)
	rootCmd.AddCommand(auth.AuthCmd)
	rootCmd.AddCommand(cfg.Cfgcmd)
	rootCmd.AddCommand(initialize())
	rootCmd.PersistentFlags().
		StringVar(&cfgFile, "config", "", "config file (default is $HOME/.DurpCLI.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	usr, err := user.Current()
	if err != nil {
		os.Exit(1)
	}
	viper.SetConfigType("yaml")
	viper.SetConfigName(".durpcli")
	viper.AddConfigPath(usr.HomeDir)
}

func loadConfig() {
	viper.AutomaticEnv()
	viper.AddConfigPath("~/.durpcli.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found. Creating a new one with defaults...")
		viper.SafeWriteConfig()
	}
}
