package dotfiles

import (
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/manifoldco/promptui"
	"strings"
)

func readUserInputBranch() {
	if branch != "" {
		return
	}
	var existingBranch string
	var allExistingBranches []string
	if repository != nil {
		head, err := repository.Head()
		if err != nil {
			logger.Fatal("Failed to get HEAD from repository: ", err)
		}
		existingBranch = head.Name().Short()
		logger.Info("Currently selected branch: ", existingBranch)

		allBranches, errAB := repository.Branches()
		if errAB != nil {
			logger.Fatal("Failed to get branches from repository: ", errAB)
		}
		errAllBranch := allBranches.ForEach(func(ref *plumbing.Reference) error {
			allExistingBranches = append(allExistingBranches, ref.Name().Short())
			return nil
		})

		if errAllBranch != nil {
			logger.Fatal("Failed to iterate branches: ", errAllBranch)
		}
	}

	prompt := promptui.Select{
		Label: "Branch",
		Items: allExistingBranches,
		Searcher: func(input string, index int) bool {
			branch := allExistingBranches[index]
			name := strings.Replace(strings.ToLower(branch), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)
			return strings.Contains(name, input)
		},
	}
	_, result, err := prompt.Run()
	if err != nil {
		utils.IsInterrupt(err)
		logger.Fatal("Prompt failed: ", err)
	}
	branch = result
}
