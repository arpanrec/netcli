package cmd

import (
	"github.com/arpanrec/netcli/internal/serverworkspace"
	"github.com/spf13/cobra"
)

func getNebulaCMD() *cobra.Command {
	var nebulaRunnerCMD = &cobra.Command{
		Use:   "nebula",
		Short: "Nebula Runner",
		Long:  "Nebula Runner is a tool to [arpanrec.nebula](https://github.com/arpanrec/arpanrec.nebula/tree/main/playbooks) playbooks",
	}
	nebulaRunnerCMD.AddCommand(serverworkspace.GetServerWorkspaceCMD())
	return nebulaRunnerCMD
}
