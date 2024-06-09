package dotfiles

import (
	"errors"
	"github.com/arpanrec/netcli/internal/logger"
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func checkout() {
	currentHeadRef, errCurrentHeadRef := repository.Head()
	if errCurrentHeadRef != nil {
		logger.Fatal("Failed to get current HEAD reference: ", errCurrentHeadRef)
	}
	logger.Info("Current HEAD target: ", currentHeadRef.Target())
	if branch == currentHeadRef.Name().Short() {
		logger.Info("Already on branch: ", branch)
		return
	}
	wt, errWt := repository.Worktree()
	if errWt != nil {
		logger.Fatal("Failed to get worktree: ", errWt)
	}
	logger.Info("Checking out branch: ", branch)
	errCheckout := wt.Checkout(&gogit.CheckoutOptions{
		Branch: plumbing.NewRemoteReferenceName("origin", branch),
		Force:  false,
		Create: false,
		Keep:   true,
	})
	if errCheckout != nil {
		logger.Fatal("Failed to checkout branch: ", errCheckout)
	}
	logger.Info("Checked out branch: ", branch)

	logger.Info("Pulling latest changes")
	errPull := wt.Pull(&gogit.PullOptions{
		RemoteName:    "origin",
		Auth:          authMethod,
		Force:         false,
		ReferenceName: plumbing.ReferenceName("refs/heads/" + branch),
	})
	if errPull != nil {
		if errors.Is(gogit.NoErrAlreadyUpToDate, errPull) {
			logger.Info("Already up to date")
		} else {
			logger.Fatal("Failed to pull latest changes: ", errPull)
		}

	}
}
