package cmd

import (
	"github.com/arpanrec/netcli/internal/dotfiles"
	"github.com/spf13/cobra"
)

var netCLI = &cobra.Command{
	Use:   "netcli",
	Short: "Set of utilities for bootstrapping a new machine",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MaximumNArgs(0)(cmd, args); err != nil {
			return err
		}
		return nil
	},
}

func Execute() error {
	return netCLI.Execute()
}

func init() {
	netCLI.PersistentFlags().BoolP("silent", "s", false, "Silent mode")
	netCLI.AddCommand(dotfiles.Cmd)
}
