package serverworkspace

import (
	"github.com/arpanrec/netcli/internal/vars"
	"github.com/spf13/cobra"
)

func GetServerWorkspaceCMD() *cobra.Command {
	var serverWorkspaceCMD = &cobra.Command{
		Use: "serverworkspace",
		Run: main,
		Long: `Setup workspace for development using

[server workspace playbook](https://github.com/arpanrec/arpanrec.nebula/blob/main/playbooks/server_workspace.md)`,
		Short: "Setup workspace for development using server workspace playbook",
		Args:  vars.IDontAllowArguments,
	}

	serverWorkspaceCMD.Flags().BoolVarP(&nodeJs, "nodejs", "", false, "Install Node.js")
	serverWorkspaceCMD.Flags().BoolVarP(&golang, "go", "", false, "Install GoLang")
	serverWorkspaceCMD.Flags().BoolVarP(&java, "java", "", false, "Install Java")
	serverWorkspaceCMD.Flags().BoolVarP(&terminal, "terminal", "", false, "Install Terminal")
	serverWorkspaceCMD.Flags().BoolVarP(&terraform, "terraform", "", false, "Install Terraform")
	serverWorkspaceCMD.Flags().BoolVarP(&vault, "vault", "", false, "Install Vault")
	serverWorkspaceCMD.Flags().BoolVarP(&pulumi, "pulumi", "", false, "Install Pulumi")
	serverWorkspaceCMD.Flags().BoolVarP(&bws, "bws", "", false, "Install BWS")
	serverWorkspaceCMD.Flags().StringVarP(&rawArgs, "raw", "", "",
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
