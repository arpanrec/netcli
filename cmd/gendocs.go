package cmd

import (
	"github.com/arpanrec/netcli/internal/constants"
	"github.com/arpanrec/netcli/internal/gendocs"
	"github.com/spf13/cobra"
)

var genDocs = &cobra.Command{
	Use:    "gendocs",
	Hidden: true,
	Short:  "Generate markdown",
	Long:   "Generate markdown documentation in the docs directory.",
	Args:   constants.IDontAllowArguments,
	Run:    gendocs.Main,
}

func addGenDocsToRoot() {
	genDocs.Flags().StringVarP(&gendocs.OutputDirectory, "output", "o", "./docs", "Directory to store the generated markdown files.")
	netCLI.AddCommand(genDocs)
}
