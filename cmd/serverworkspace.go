package cmd

import (
	"github.com/arpanrec/netcli/internal/serverworkspace"
	"github.com/spf13/cobra"
)

func getServerWorkspaceCMD() *cobra.Command {
	var serverWorkspaceCMD = &cobra.Command{
		Use:   "serverworkspace",
		Run:   serverworkspace.Main,
		Long:  `Setup workspace for development using [server workspace playbook](https://github.com/arpanrec/arpanrec.nebula/blob/main/playbooks/server_workspace.md)`,
		Short: "Setup workspace for development using server workspace playbook",
	}

	serverWorkspaceCMD.Flags().BoolVarP(&serverworkspace.NodeJs, "nodejs", "", false, "Install Node.js")
	serverWorkspaceCMD.Flags().BoolVarP(&serverworkspace.Go, "go", "", false, "Install GoLang")
	serverWorkspaceCMD.Flags().BoolVarP(&serverworkspace.Java, "java", "", false, "Install Java")
	serverWorkspaceCMD.Flags().BoolVarP(&serverworkspace.Terminal, "terminal", "", false, "Install Terminal")
	serverWorkspaceCMD.Flags().BoolVarP(&serverworkspace.Terraform, "terraform", "", false, "Install Terraform")
	serverWorkspaceCMD.Flags().BoolVarP(&serverworkspace.Vault, "vault", "", false, "Install Vault")
	serverWorkspaceCMD.Flags().BoolVarP(&serverworkspace.Pulumi, "pulumi", "", false, "Install Pulumi")
	serverWorkspaceCMD.Flags().BoolVarP(&serverworkspace.BWS, "bws", "", false, "Install BWS")
	serverWorkspaceCMD.Flags().StringVarP(&serverworkspace.RawArgs, "raw", "", "",
		"Pass raw arguments to the script. Example: --raw \"--nodejs --go --java\", this will also add the local config file: "+serverworkspace.LocalConfigPath)
	serverWorkspaceCMD.MarkFlagsMutuallyExclusive("raw", "nodejs")
	serverWorkspaceCMD.MarkFlagsMutuallyExclusive("raw", "go")
	serverWorkspaceCMD.MarkFlagsMutuallyExclusive("raw", "java")
	serverWorkspaceCMD.MarkFlagsMutuallyExclusive("raw", "terminal")
	serverWorkspaceCMD.MarkFlagsMutuallyExclusive("raw", "terraform")
	serverWorkspaceCMD.MarkFlagsMutuallyExclusive("raw", "vault")
	serverWorkspaceCMD.MarkFlagsMutuallyExclusive("raw", "pulumi")
	serverWorkspaceCMD.MarkFlagsMutuallyExclusive("raw", "bws")

	return serverWorkspaceCMD
}
