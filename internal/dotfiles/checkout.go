package dotfiles

import (
	"fmt"
	"github.com/arpanrec/netcli/internal/logger"
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"os"
)

func checkout() {
	rConfig, _ := repository.Config()
	fmt.Println(rConfig.Raw.Section("status").Option("showUntrackedFiles"))

	ref, _ := repository.Head()
	fmt.Println(ref.Name())

	references, err := repository.References()
	if err != nil {
		logger.Fatal("Failed to get references: ", err)
	}

	_ = references.ForEach(func(ref *plumbing.Reference) error {
		fmt.Println(ref.Name())
		return nil
	})

	workTree, errWorkTree := repository.Worktree()
	if errWorkTree != nil {
		logger.Fatal("Failed to get worktree: ", errWorkTree)
	}

	pullErr := workTree.Pull(&gogit.PullOptions{
		Auth:     authMethod,
		Progress: os.Stdout,
	})
	if pullErr != nil {
		if pullErr.Error() == "already up-to-date" {
			logger.Info("Repository is already up to date")
		} else {
			logger.Fatal("Failed to pull repository: ", pullErr)
		}
	}
	xx, _ := workTree.Status()
	fmt.Println(xx.IsUntracked(".prettierrc.mjs"))

	logger.Info("Checking out branch: ", branch)
	errCheckout := workTree.Checkout(&gogit.CheckoutOptions{
		Branch: "refs/heads/office",
		Keep:   true,
		Force:  false,
	})
	if errCheckout != nil {
		logger.Fatal("Failed to checkout branch: ", errCheckout)
	}
	// err = workTree.Reset(&gogit.ResetOptions{
	// 	Mode: gogit.HardReset,
	// })
	// if err != nil {
	// 	logger.Fatal("Failed to reset worktree: ", err)
	// }
}
