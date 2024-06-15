package serverworkspace

import (
	"errors"
	"github.com/arpanrec/netcli/internal/utils"
)

func askForConfirmation() {
	if isSilent {
		return
	}
	// utils.PromptBool("Do you want to install the following packages?", true)

	if RawArgs == "" && !nodeJsProvided && !goProvided && !javaProvided &&
		!terminalProvided && !terraformProvided && !vaultProvided && !pulumiProvided && !bwsProvided {

		wantRawArgs := utils.PromptBool("Do you want to provide raw arguments", false)

		if wantRawArgs {
			RawArgs = utils.PromptString("Enter raw arguments", "", func(input string) error {
				if input == "" {
					return errors.New("raw arguments cannot be empty")
				}
				return nil
			})
			return
		}
	}

	if !NodeJs && !nodeJsProvided {
		NodeJs = utils.PromptBool("Install NodeJs", false)
	}

	if !Go && !goProvided {
		Go = utils.PromptBool("Install Go", false)
	}

	if !Java && !javaProvided {
		Java = utils.PromptBool("Install Java", false)
	}

	if !Terminal && !terminalProvided {
		Terminal = utils.PromptBool("Install Terminal", false)
	}

	if !Terraform && !terraformProvided {
		Terraform = utils.PromptBool("Install Terraform", false)
	}

	if !Vault && !vaultProvided {
		Vault = utils.PromptBool("Install Vault", false)
	}

	if !Pulumi && !pulumiProvided {
		Pulumi = utils.PromptBool("Install Pulumi", false)
	}

	if !BWS && !bwsProvided {
		BWS = utils.PromptBool("Install BWS", false)
	}
}
