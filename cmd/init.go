package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/DeveloperDurp/DurpCLI/internal/tui"
)

func initialize() *cobra.Command {
	init := &cobra.Command{
		Use:     "initialize",
		Short:   "init the rcl cfg.",
		Long:    "init provision the rcl configuration file.",
		Example: "rkl init",
		Aliases: []string{"i", "init"},
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			// Environment variables are expected to be ALL CAPS
			viper.AutomaticEnv()
			viper.SetEnvPrefix("rkl")
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			p := tea.NewProgram(tui.InitialModel())
			if _, err := p.Run(); err != nil {
				fmt.Printf("Alas, there's been an error: %v", err)
				os.Exit(1)
			}
			return nil
		},
	}
	return init
}
