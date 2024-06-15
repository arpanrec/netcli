package serverworkspace

import (
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/spf13/cobra"
	"strconv"
)

func Main(cmd *cobra.Command, _ []string) {
	isS, err := strconv.ParseBool(cmd.Flag("silent").Value.String())
	if err != nil {
		logger.Fatal("Failed to get silent flag", err)
	}
	isSilent = isS
	NodeJsProvided = cmd.Flag("nodejs").Changed
	GoProvided = cmd.Flag("go").Changed
	JavaProvided = cmd.Flag("java").Changed
	TerminalProvided = cmd.Flag("terminal").Changed
	TerraformProvided = cmd.Flag("terraform").Changed
	VaultProvided = cmd.Flag("vault").Changed
	PulumiProvided = cmd.Flag("pulumi").Changed
	BWSProvided = cmd.Flag("bws").Changed
	RawArgsProvided = cmd.Flag("raw").Changed
	preCheck()
	askForConfirmation()
	createVenv()
	ifLocalConfigExist()
	writeInventoryFile()
	run()
}
