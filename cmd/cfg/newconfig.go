package cfg

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var newcfgcmd = &cobra.Command{
	Use:   "newconfig",
	Short: "Generates a config file using current config",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.WriteConfigAs(".DurpCLI.yaml")
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {

	Cfgcmd.AddCommand(newcfgcmd)
}
