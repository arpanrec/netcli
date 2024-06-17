package dotfiles

import (
	"strings"

	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/manifoldco/promptui"
)

func readUserInputBranch() {
	if Branch != "" {
		return
	}
	var currentLocalBranch string
	allExistingBranches := make([]string, 0)
	setLocalBranches(&currentLocalBranch, &allExistingBranches)

	for _, ref := range remoteRefs {
		if !ref.Name().IsBranch() {
			continue
		}
		shortName := ref.Name().Short()
		if utils.IfElementInSlice(&allExistingBranches, &shortName) == -1 {
			allExistingBranches = append(allExistingBranches, ref.Name().Short())
		}
	}

	defaultRef := getDefaultRemoteBranch()

	if isSilent {
		if currentLocalBranch != "" {
			Branch = currentLocalBranch
			logger.Info("Using existing branch: ", Branch)
		} else {
			Branch = defaultRef.Name().Short()
			logger.Info("Using HEAD target branch: ", Branch)
		}
		return
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
		utils.IsInterrupt(&err)
		logger.Fatal("Prompt failed: ", err)
	}
	Branch = result
}

func setLocalBranches(currentLocalBranch *string, allExistingBranches *[]string) {
	if repository != nil {
		head, err := repository.Head()
		if err != nil {
			logger.Fatal("Failed to get HEAD from local repository: ", err)
		}
		*currentLocalBranch = head.Name().Short()
		logger.Info("Current local tracking branch: ", *currentLocalBranch)

		allBranches, errAB := repository.Branches()
		if errAB != nil {
			logger.Fatal("Failed to get all branches from repository: ", errAB)
		}
		errAllBranch := allBranches.ForEach(func(ref *plumbing.Reference) error {
			if ref.Name().IsBranch() {
				*allExistingBranches = append(*allExistingBranches, ref.Name().Short())
			}
			return nil
		})

		if errAllBranch != nil {
			logger.Fatal("Failed to iterate branches: ", errAllBranch)
		}
	}
}

func getDefaultRemoteBranch() *plumbing.Reference {
	var headRefTargetShort string
	var defaultRef *plumbing.Reference
	if remoteRefs == nil {
		logger.Fatal("No remote branches found")
	}
	for _, ref := range remoteRefs {
		if ref.Name().Short() == "HEAD" {
			headRefTargetShort = ref.Target().Short()
			break
		}
	}
	if headRefTargetShort == "" {
		logger.Fatal("HEAD target not found")
	}

	for _, ref := range remoteRefs {
		sortName := ref.Name().Short()
		if !ref.Name().IsBranch() {
			continue
		}
		if sortName == headRefTargetShort {
			defaultRef = ref
		}
	}
	if defaultRef != nil {
		logger.Debug("Default remote branch: ", defaultRef.Name().Short())
	} else {
		logger.Fatal("Default remote branch not found")
	}
	return defaultRef
}
