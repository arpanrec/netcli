package serverworkspace

import (
	"github.com/spf13/cobra"
)

func Main(cmd *cobra.Command, _ []string) {

	NodeJsProvided = cmd.Flag("nodejs").Changed
	GoProvided = cmd.Flag("go").Changed
	JavaProvided = cmd.Flag("java").Changed
	TerminalProvided = cmd.Flag("terminal").Changed
	TerraformProvided = cmd.Flag("terraform").Changed
	VaultProvided = cmd.Flag("vault").Changed
	PulumiProvided = cmd.Flag("pulumi").Changed
	BWSProvided = cmd.Flag("bws").Changed
	createVenv()
}
