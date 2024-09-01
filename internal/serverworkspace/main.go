package serverworkspace

import (
	"strconv"

	"github.com/arpanrec/netcli/internal/logger"
	"github.com/spf13/cobra"
)

func main(cmd *cobra.Command, _ []string) {
	isS, err := strconv.ParseBool(cmd.Flag("silent").Value.String())
	if err != nil {
		logger.Fatal("Failed to get silent flag", err)
	}
	isSilent = isS
	nodeJsProvided = cmd.Flag("nodejs").Changed
	goProvided = cmd.Flag("go").Changed
	javaProvided = cmd.Flag("java").Changed
	terminalProvided = cmd.Flag("terminal").Changed
	terraformProvided = cmd.Flag("terraform").Changed
	vaultProvided = cmd.Flag("vault").Changed
	pulumiProvided = cmd.Flag("pulumi").Changed
	bwsProvided = cmd.Flag("bws").Changed
	rawArgsProvided = cmd.Flag("raw").Changed
	preCheck()
	askForConfirmation()
	createVenv()
	ifLocalConfigExist()
	writeInventoryFile()
	run()
}
