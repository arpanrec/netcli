package dotfiles

import (
	"github.com/arpanrec/netcli/internal/logger"
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
)

var repositoryUrl string

var branch string

var directory string

var isCleanInstall bool

var isResetHead bool

var repository *gogit.Repository

var remote *gogit.Remote

var sshKeyPath string

var sshKeyPassphrase string

var authMethod transport.AuthMethod

var remoteRefs []*plumbing.Reference

func install(cmd *cobra.Command, _ []string) {
	isSilent, err := strconv.ParseBool(cmd.Flag("silent").Value.String())
	if err != nil {
		logger.Fatal("Failed to get silent flag", err)
	}
	logger.Debug("Install called with silent: ", isSilent)
	logger.Debug("Repository from flag: ", repositoryUrl)
	logger.Debug("Branch from flag: ", branch)
	logger.Debug("Directory from flag: ", directory)
	logger.Debug("Clean install flag: ", isCleanInstall)
	logger.Debug("Reset HEAD flag: ", isResetHead)

	if isSilent {
		if repositoryUrl == "" {
			logger.Fatal("Repository is not provided, but running in silent mode")
		}
		if directory == "" {
			logger.Fatal("Directory is not provided, but running in silent mode")
		}
		if branch == "" {
			logger.Fatal("Branch is not provided, but running in silent mode")
		}
	}
	readUserInputDirectory()
	validateDirectoryAndLoadRepo()
	logger.Info("Directory: ", directory)

	readUserInputRepositoryUrl()
	validateRepositoryUrl()
	logger.Info("Repository: ", repositoryUrl)

	createRemote()

	readUserInputBranch()
	logger.Info("Branch: ", branch)

}

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
	} else {
		var authMethod transport.AuthMethod
		if strings.HasPrefix(repositoryUrl, "git@") {
			defaultUserSettings := ssh.DefaultSSHConfig.Get("github.com", "IdentityFile")
			logger.Debug("Default user settings: ", defaultUserSettings)
			am, errAuth := ssh.NewPublicKeysFromFile("git", os.Getenv("HOME")+"/.ssh/id_rsa", "")
			if errAuth != nil {
				logger.Fatal("Failed to create SSH agent auth: ", errAuth)
			}
			authMethod = am
			logger.Debug("Using SSH auth method")
		}

		rem := gogit.NewRemote(memory.NewStorage(), &config.RemoteConfig{
			Name: "origin",
			URLs: []string{repositoryUrl},
		})

		refs, errRefs := rem.List(&gogit.ListOptions{
			Auth: authMethod,
		})

		if errRefs != nil {
			logger.Fatal("Failed to get branches from remote: ", errRefs)
		}

		for _, ref := range refs {
			allExistingBranches = append(allExistingBranches, ref.Name().Short())
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
		logger.Fatal("Prompt failed: ", err)
	}
	branch = result
}
