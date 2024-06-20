package dotfiles

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/manifoldco/promptui"
)

func _() { // func checkout() {} // Because I don't have a degree in plumbing.
	currentHeadRef, errCurrentHeadRef := repository.Head()
	if errCurrentHeadRef != nil {
		logger.Fatal("Failed to get current HEAD reference: ", errCurrentHeadRef)
	}
	logger.Info("Current HEAD target: ", currentHeadRef.Name().Short())
	if Branch == currentHeadRef.Name().Short() { // Bug
		logger.Info("Already on branch: ", Branch)
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
		if sortName == Branch || sortName == "origin/"+Branch {
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
		logger.Fatal("Branch not found: ", Branch)
	}

	wt, errWt := repository.Worktree()
	if errWt != nil {
		logger.Fatal("Failed to get worktree: ", errWt)
	}
	logger.Info("Checking out branch: ", Branch)
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
	logger.Info("Checked out branch: ", Branch)

	logger.Info("Pulling latest changes")
	errPull := wt.Pull(&gogit.PullOptions{
		RemoteName:    "origin",
		Auth:          authMethod,
		Force:         false,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s:refs/heads/%s", Branch, Branch)),
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
	logger.Info("Checking out branch: ", Branch)
	resetHead()
	cmd := fmt.Sprintf("checkout %s", Branch)
	logger.Info("Executing checkout command: ", cmd)
	out := gitExec(&cmd)
	logger.Info("Checkout command output: ", out)
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
			IsResetHead = true
		}
	}
	if !IsResetHead {
		return
	}
	logger.Info("Resetting HEAD")
	cmd := "reset --hard HEAD"
	out := gitExec(&cmd)
	logger.Info("Reset command output: ", out)
}

func addToRc() {
	logger.Info("Adding alias to rc file")
	aliasesEntry := fmt.Sprintf("alias dotfiles=\"'git --git-dir=%s --work-tree=%s'\"", GitDirectory, workTreeDir)
	logger.Info("Adding alias to rc file" + aliasesEntry)
	files := []string{".bashrc", ".zshrc", ".bash_profile", ".profile", ".bash_aliases", ".aliasrc"}
	for _, file := range files {
		rcFile := path.Join(workTreeDir, file)
		cmd := fmt.Sprintf("echo '%s' | tee -a %s", aliasesEntry, rcFile)
		_, err := utils.BashExec(&cmd)
		if err != nil {
			logger.Warn("Failed to add alias to rc file: ", rcFile)
		}
		logger.Info("Added alias to rc file: ", rcFile)
	}
}
