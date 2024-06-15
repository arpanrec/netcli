package cmd

import (
	"github.com/arpanrec/netcli/internal/serverworkspace"
	"github.com/spf13/cobra"
)

func getServerWorkspaceCMD() *cobra.Command {
	var serverWorkspaceCMD = &cobra.Command{
		Use: "serverworkspace",
		Run: serverworkspace.Main,
	}

	serverWorkspaceCMD.Flags().BoolVarP(&serverworkspace.NodeJs, "nodejs", "", false, "Install Node.js")
	serverWorkspaceCMD.Flags().BoolVarP(&serverworkspace.Go, "go", "", false, "Install GoLang")
	serverWorkspaceCMD.Flags().BoolVarP(&serverworkspace.Java, "java", "", false, "Install Java")
	serverWorkspaceCMD.Flags().BoolVarP(&serverworkspace.Terminal, "terminal", "", false, "Install Terminal")
	serverWorkspaceCMD.Flags().BoolVarP(&serverworkspace.Terraform, "terraform", "", false, "Install Terraform")
	serverWorkspaceCMD.Flags().BoolVarP(&serverworkspace.Vault, "vault", "", false, "Install Vault")
	serverWorkspaceCMD.Flags().BoolVarP(&serverworkspace.Pulumi, "pulumi", "", false, "Install Pulumi")
	serverWorkspaceCMD.Flags().BoolVarP(&serverworkspace.BWS, "bws", "", false, "Install BWS")

	return serverWorkspaceCMD
}
