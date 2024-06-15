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

	if RawArgs == "" && !NodeJsProvided && !GoProvided && !JavaProvided &&
		!TerminalProvided && !TerraformProvided && !VaultProvided && !PulumiProvided && !BWSProvided {

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

	if !NodeJs && !NodeJsProvided {
		NodeJs = utils.PromptBool("Install NodeJs", false)
	}
	if !Go && !GoProvided {
		Go = utils.PromptBool("Install Go", false)
	}
	if !Java && !JavaProvided {
		Java = utils.PromptBool("Install Java", false)
	}

	if !Terminal && !TerminalProvided {
		Terminal = utils.PromptBool("Install Terminal", false)
	}

	if !Terraform && !TerraformProvided {
		Terraform = utils.PromptBool("Install Terraform", false)
	}

	if !Vault && !VaultProvided {
		Vault = utils.PromptBool("Install Vault", false)
	}

	if !Pulumi && !PulumiProvided {
		Pulumi = utils.PromptBool("Install Pulumi", false)
	}

	if !BWS && !BWSProvided {
		BWS = utils.PromptBool("Install BWS", false)
	}
}
