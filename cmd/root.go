package cmd

import (
	"errors"
	"github.com/arpanrec/netcli/internal/constants"
	"github.com/arpanrec/netcli/internal/dotfiles"
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"log"
)

var netCLI = &cobra.Command{
	Use:     constants.NetCliUse,
	Short:   constants.NetCliShort,
	Long:    constants.NetCliLong,
	Version: constants.NetCliVersion,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MaximumNArgs(0)(cmd, args); err != nil {
			return errors.New("No arguments are allowed. Error: " + err.Error())
		}
		return nil
	},
}

func Execute() error {
	err := doc.GenMarkdownTree(netCLI, "./docs")
	if err != nil {
		log.Fatal(err)
	}
	return netCLI.Execute()
}

func init() {
	netCLI.PersistentFlags().BoolP("silent", "s", false, "Silent mode")
	netCLI.PersistentFlags().BoolVarP(&logger.DebugMode, "debug-logging", "", false,
		"Enable debug logging")
	netCLI.AddCommand(dotfiles.Cmd)
}
