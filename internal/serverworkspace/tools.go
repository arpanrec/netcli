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

	if rawArgs == "" && !nodeJsProvided && !goProvided && !javaProvided &&
		!terminalProvided && !terraformProvided && !vaultProvided && !pulumiProvided && !bwsProvided &&
		!bitwardenDesktopProvided && !mattermostDesktopProvided && !telegramDesktopProvided &&
		!postmanProvided && !codeProvided && !themesProvided {

		wantRawArgs := utils.PromptBool("Do you want to provide raw arguments", false)

		if wantRawArgs {
			rawArgs = utils.PromptString("Enter raw arguments", "", func(input string) error {
				if input == "" {
					return errors.New("raw arguments cannot be empty")
				}
				return nil
			})
			return
		}
	}

	if rawArgs != "" {
		return
	}

	if !nodeJs && !nodeJsProvided {
		nodeJs = utils.PromptBool("Install NodeJs", false)
	}

	if !golang && !goProvided {
		golang = utils.PromptBool("Install Go", false)
	}

	if !java && !javaProvided {
		java = utils.PromptBool("Install Java", false)
	}

	if !terminal && !terminalProvided {
		terminal = utils.PromptBool("Install Terminal", false)
	}

	if !terraform && !terraformProvided {
		terraform = utils.PromptBool("Install Terraform", false)
	}

	if !vault && !vaultProvided {
		vault = utils.PromptBool("Install Vault", false)
	}

	if !pulumi && !pulumiProvided {
		pulumi = utils.PromptBool("Install Pulumi", false)
	}

	if !bws && !bwsProvided {
		bws = utils.PromptBool("Install BWS", false)
	}

	if !bitwardenDesktop && !bitwardenDesktopProvided {
		bitwardenDesktop = utils.PromptBool("Install Bitwarden Desktop", false)
	}

	if !mattermostDesktop && !mattermostDesktopProvided {
		mattermostDesktop = utils.PromptBool("Install Mattermost Desktop", false)
	}

	if !telegramDesktop && !telegramDesktopProvided {
		telegramDesktop = utils.PromptBool("Install Telegram Desktop", false)
	}

	if !postman && !postmanProvided {
		postman = utils.PromptBool("Install Postman", false)
	}

	if !code && !codeProvided {
		code = utils.PromptBool("Install Visual Studio Code", false)
	}

	if !themes && !themesProvided {
		themes = utils.PromptBool("Install Themes", false)
	}
}
