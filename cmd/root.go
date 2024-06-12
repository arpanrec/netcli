package cmd

import (
	"errors"

	"github.com/arpanrec/netcli/internal/constants"
	"github.com/arpanrec/netcli/internal/dotfiles"
	"github.com/arpanrec/netcli/internal/gendocs"
	"github.com/spf13/cobra"
)

var netCLI = &cobra.Command{
	Use:     "netcli",
	Short:   constants.NetCliShort,
	Long:    constants.NetCliLong,
	Example: "netcli dotfiles -h",
	Version: constants.Version,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MaximumNArgs(0)(cmd, args); err != nil {
			return errors.New("No arguments are allowed. Error: " + err.Error())
		}
		return nil
	},
}

func Execute() error {
	return netCLI.Execute()
}

func init() {
	netCLI.PersistentFlags().BoolP("silent", "s", false, "Silent mode")

	// Just for documentation not actually used.
	// Actual logging is done in internal/logger/logger.go
	netCLI.PersistentFlags().BoolP("debug-logging", "", false,
		"Enable debug logging. This can be set using the environment variable DEBUG=true.")

	netCLI.AddCommand(dotfiles.Cmd)
	netCLI.AddCommand(gendocs.Cmd)
}
