package gendocs

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func Main(cmd *cobra.Command, _ []string) {
	outputDirectoryProvided = cmd.Flag("output").Changed
	preChecks()
	outputDir()
	rootCmd := cmd.Root()
	rootCmd.DisableAutoGenTag = true
	err := doc.GenMarkdownTree(rootCmd, OutputDirectory)
	if err != nil {
		log.Fatal("error generating markdown documentation" + err.Error())
	}
	createReadme()
	createWebRunScript()
}
