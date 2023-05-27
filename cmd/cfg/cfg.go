package cfg

import (
	"github.com/spf13/cobra"
)

var Cfgcmd = &cobra.Command{
	Use:   "cfg",
	Short: "All things Authorization",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
