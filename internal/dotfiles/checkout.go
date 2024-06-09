package dotfiles

import (
	"errors"
	"fmt"
	"github.com/arpanrec/netcli/internal/logger"
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"os"
)

func _() { // func checkout() {}
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

func checkoutWithCmd() {
	logger.Info("Checking out branch: ", branch)
	cmd := fmt.Sprintf("checkout %s", branch)
	logger.Info("Executing checkout command: ", cmd)
	out := gitExec(&cmd)
	logger.Info("Checkout command output: ", out)
}
