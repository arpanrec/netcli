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

func init() {
	genDocs.Flags().StringVarP(&gendocs.OutputDirectory, "output", "o", "./docs", "output directory")
}