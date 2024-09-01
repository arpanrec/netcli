package cmd

import (
	"github.com/arpanrec/netcli/internal/gendocs"
	"github.com/arpanrec/netcli/internal/vars"
	"github.com/spf13/cobra"
)

func getGenDocsCMD() *cobra.Command {
	var genDocs = &cobra.Command{
		Use:    "gendocs",
		Hidden: true,
		Short:  "Generate markdown",
		Long:   "Generate markdown documentation in the docs directory.",
		Args:   vars.IDontAllowArguments,
		Run:    gendocs.Main,
	}
	genDocs.Flags().StringVarP(&gendocs.OutputDirectory, "output", "o", "./docs", "Directory to store the generated markdown files.")

	return genDocs
}
