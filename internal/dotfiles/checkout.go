package dotfiles

import (
	"errors"
	"fmt"
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/manifoldco/promptui"
	"os"
	"strings"
)

func _() { // func checkout() {} // Because I don't have a degree in plumbing.
	currentHeadRef, errCurrentHeadRef := repository.Head()
	if errCurrentHeadRef != nil {
		logger.Fatal("Failed to get current HEAD reference: ", errCurrentHeadRef)
	}
	logger.Info("Current HEAD target: ", currentHeadRef.Name().Short())
	if branch == currentHeadRef.Name().Short() { // Bug
		logger.Info("Already on branch: ", branch)
		return
	}

	references, errRef := repository.Storer.IterReferences()
	if errRef != nil {
		logger.Fatal("Failed to get references: ", errRef)
	}

	var remoteRef *plumbing.Reference
	var localRef *plumbing.Reference
	err := references.ForEach(func(ref *plumbing.Reference) error {
		sortName := ref.Name().Short()
		if sortName == branch || sortName == "origin/"+branch {
			logger.Info("Branch found: ", sortName)
			if ref.Name().IsRemote() {
				remoteRef = ref
			} else {
				localRef = ref
			}
		}
		return nil
	})
	if err != nil {
		logger.Fatal("Failed to iterate references: ", err)
	}

	if remoteRef == nil && localRef == nil {
		logger.Fatal("Branch not found: ", branch)
	}

	wt, errWt := repository.Worktree()
	if errWt != nil {
		logger.Fatal("Failed to get worktree: ", errWt)
	}
	logger.Info("Checking out branch: ", branch)
	branchCoOpts := gogit.CheckoutOptions{
		Branch: remoteRef.Name(),
		Force:  false,
		Create: true,
		Keep:   true,
	}
	errCheckout := wt.Checkout(&branchCoOpts)
	if errCheckout != nil {
		logger.Fatal("Failed to checkout branch: ", errCheckout)
	}
	logger.Info("Checked out branch: ", branch)

	logger.Info("Pulling latest changes")
	errPull := wt.Pull(&gogit.PullOptions{
		RemoteName:    "origin",
		Auth:          authMethod,
		Force:         false,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s:refs/heads/%s", branch, branch)),
		Depth:         0,
		SingleBranch:  true,
		Progress:      os.Stdout,
	})
	if errPull != nil {
		if errors.Is(gogit.NoErrAlreadyUpToDate, errPull) {
			logger.Info("Already up to date")
		} else {
			logger.Fatal("Failed to pull latest changes: ", errPull)
		}
	}
}

func checkout() {
	logger.Info("Checking out branch: ", branch)
	resetHead()
	cmd := fmt.Sprintf("checkout %s", branch)
	logger.Info("Executing checkout command: ", cmd)
	out := gitExec(&cmd)
	logger.Info("Checkout command output: ", out)
	resetHead()
}

func resetHead() {
	if !isSilent && !isResetHeadProvided {
		logger.Info("Do you want reset to the HEAD?")
		logger.Info("This will discard all changes and reset to the latest commit.")
		options := []string{"No", "Yes"}
		prompt := promptui.Select{
			Label: "Reset to HEAD?",
			Items: options,
			Searcher: func(input string, index int) bool {
				name := strings.Replace(strings.ToLower(options[index]), " ", "", -1)
				input = strings.Replace(strings.ToLower(input), " ", "", -1)
				return strings.Contains(name, input)
			},
		}
		_, result, err := prompt.Run()
		if err != nil {
			utils.IsInterrupt(&err)
			logger.Fatal("Prompt failed: ", err)
		}
		if result == "Yes" {
			isResetHead = true
		}
	}
	if !isResetHead {
		return
	}
	logger.Info("Resetting HEAD")
	cmd := fmt.Sprintf("reset --hard HEAD")
	out := gitExec(&cmd)
	logger.Info("Reset command output: ", out)
}
