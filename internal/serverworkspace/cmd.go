package serverworkspace

import (
	"github.com/arpanrec/netcli/internal/constants"
	"github.com/spf13/cobra"
)

func GetServerWorkspaceCMD() *cobra.Command {
	var serverWorkspaceCMD = &cobra.Command{
		Use: "serverworkspace",
		Run: main,
		Long: `Setup workspace for development using

[server workspace playbook](https://github.com/arpanrec/arpanrec.nebula/blob/main/playbooks/server_workspace.md)`,
		Short: "Setup workspace for development using server workspace playbook",
		Args:  constants.IDontAllowArguments,
	}

	serverWorkspaceCMD.Flags().BoolVarP(&NodeJs, "nodejs", "", false, "Install Node.js")
	serverWorkspaceCMD.Flags().BoolVarP(&Go, "go", "", false, "Install GoLang")
	serverWorkspaceCMD.Flags().BoolVarP(&Java, "java", "", false, "Install Java")
	serverWorkspaceCMD.Flags().BoolVarP(&Terminal, "terminal", "", false, "Install Terminal")
	serverWorkspaceCMD.Flags().BoolVarP(&Terraform, "terraform", "", false, "Install Terraform")
	serverWorkspaceCMD.Flags().BoolVarP(&Vault, "vault", "", false, "Install Vault")
	serverWorkspaceCMD.Flags().BoolVarP(&Pulumi, "pulumi", "", false, "Install Pulumi")
	serverWorkspaceCMD.Flags().BoolVarP(&BWS, "bws", "", false, "Install BWS")
	serverWorkspaceCMD.Flags().StringVarP(&RawArgs, "raw", "", "",
		"Pass raw arguments to the script. Example: --raw \"--nodejs --go --java\", this will also add the local config file: "+LocalConfigPath)
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
