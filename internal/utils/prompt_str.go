package utils

import (
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/manifoldco/promptui"
)

func PromptString(label, defaultValue string, validate func(string) error) string {
	prompt := promptui.Prompt{
		Label:    label,
		Default:  defaultValue,
		Validate: validate,
	}
	result, err := prompt.Run()
	if err != nil {
		IsInterrupt(&err)
		logger.Fatal("Prompt failed: ", err)
	}
	return result
}
