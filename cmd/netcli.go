package cmd

import (
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/vars"
	"github.com/spf13/cobra"
)

func Execute() {
	var netCLI = &cobra.Command{
		Use:     "netcli",
		Short:   vars.NetCliShort,
		Long:    vars.NetCliLong,
		Example: "netcli -h",
		Version: vars.Version(),
		Args:    vars.IDontAllowArguments,
	}
	netCLI.PersistentFlags().BoolP("silent", "s", false, "Silent mode. Do not prompt for any input.")

	// Just for documentation not actually used. Actual logging is done in internal/logger/logger.go
	netCLI.PersistentFlags().BoolP("debug-logging", "", false,
		"Enable debug logging. This can be set using the environment variable DEBUG=true.")

	// Just for documentation not actually used. Actual version is coming from cobra root netcli command.
	netCLI.Flags().BoolP("version", "v", false,
		"Print the version of the netcli command.")

	netCLI.AddCommand(getDotFilesCmd())
	netCLI.AddCommand(getGenDocsCMD())
	netCLI.AddCommand(getNebulaCMD())

	err := netCLI.Execute()
	if err != nil {
		logger.Fatal("Failed to execute netcli command", err)
	}
}
