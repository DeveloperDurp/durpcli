package cfg

import (
	"os"
	"os/user"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var newcfgcmd = &cobra.Command{
	Use:   "newconfig",
	Short: "Generates a config file using current config",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		usr, err := user.Current()
		if err != nil {
			os.Exit(1)
		}
		viper.SetConfigType("yaml")
		viper.SetConfigName(".durpcli")
		viper.AddConfigPath(usr.HomeDir)
		viper.WriteConfig()

	},
}

func init() {

	Cfgcmd.AddCommand(newcfgcmd)
}
