package utils

import (
	"strconv"
	"strings"

	"github.com/arpanrec/netcli/internal/logger"
	"github.com/manifoldco/promptui"
)

func PromptBool(label string, defaultBool bool) bool {
	options := []string{strconv.FormatBool(defaultBool), strconv.FormatBool(!defaultBool)}
	prompt := promptui.Select{
		Label: label,
		Items: options,
		Searcher: func(input string, index int) bool {
			name := strings.Replace(strings.ToLower(options[index]), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)
			return strings.Contains(name, input)
		},
	}
	_, result, err := prompt.Run()
	if err != nil {
		IsInterrupt(&err)
		logger.Fatal("Prompt failed: ", err)
	}
	res, errParsebool := strconv.ParseBool(result)
	if errParsebool != nil {
		logger.Fatal("Failed to parse bool: ", errParsebool)
	}
	return res
}
