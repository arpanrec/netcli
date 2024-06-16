package cmd

import (
	"github.com/arpanrec/netcli/internal/serverworkspace"
	"github.com/spf13/cobra"
)

func getNebulaCMD() *cobra.Command {
	var nebulaRunnerCMD = &cobra.Command{
		Use: "nebula",
	}
	nebulaRunnerCMD.AddCommand(serverworkspace.GetServerWorkspaceCMD())
	return nebulaRunnerCMD
}
