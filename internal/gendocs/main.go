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
		rootCmd := cmd.Root()
		rootCmd.DisableAutoGenTag = true
		err := doc.GenMarkdownTree(rootCmd, "./docs")
		if err != nil {
			log.Fatal(err)
		}
	},
}
