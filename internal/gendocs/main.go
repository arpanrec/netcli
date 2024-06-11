package gendocs

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"log"
)

var Cmd = &cobra.Command{
	Use:   "gendocs",
	Short: "Generate markdown",
	Long:  "Generate markdown documentation in the docs directory.",
	Run: func(cmd *cobra.Command, args []string) {
		outputDirectoryProvided = cmd.Flag("output").Changed
		preChecks()
		outputDir()
		rootCmd := cmd.Root()
		rootCmd.DisableAutoGenTag = true
		err := doc.GenMarkdownTree(rootCmd, outputDirectory)
		if err != nil {
			log.Fatal("error generating markdown documentation" + err.Error())
		}
		createReadme()
	},
}

func init() {
	Cmd.Flags().StringVarP(&outputDirectory, "output", "o", "./docs", "output directory")
}
