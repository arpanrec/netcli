package vars

import (
	"errors"

	"github.com/spf13/cobra"
)

const (
	NetCliShort = `Few utilities for bootstrapping a new machine`
	NetCliLong  = `NetCLI is a set of utilities for my day-to-day work.

This helps simplify the process of setting up a new machine, installing the necessary tools, and configuring them, etc. etc.`
)

func IDontAllowArguments(cmd *cobra.Command, args []string) error {
	if err := cobra.MaximumNArgs(0)(cmd, args); err != nil {
		return errors.New("No arguments are allowed. Error: " + err.Error())
	}
	return nil
}
