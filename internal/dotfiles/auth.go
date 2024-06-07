package dotfiles

import (
	"errors"
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
	giturl "github.com/chainguard-dev/git-urls"
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/manifoldco/promptui"
	"os"
)

func createRemote() {
	remote = gogit.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{repositoryUrl},
	})

	gitURL, errUrlParse := giturl.Parse(repositoryUrl)
	if errUrlParse != nil {
		logger.Fatal("Failed to parse repository URL: ", errUrlParse)
	}

	schema := gitURL.Scheme
	hostname := gitURL.Hostname()

	if schema == "ssh" {
		user := gitURL.User.Username()
		logger.Debug("Using SSH auth method")

		tryWithUserProvidedKey(user)

		if tryNewSSHAgentAuth(user) {
			return
		}

		if tryHostNameKeyConfig(hostname, user) {
			return
		}
	}
}

func tryWithUserProvidedKey(u string) {

	if sshKeyPath == "" && !isSilent {
		prompt := promptui.Prompt{
			Label:     "SSH Key Path",
			AllowEdit: true,
			Validate: func(s string) error {
				if s == "" {
					return nil
				}
				errAbsPath := utils.AbsPath(&s)
				if errAbsPath != nil {
					return errors.New("failed to get absolute path of SSH key, " + errAbsPath.Error())
				}
				stat, err := os.Stat(s)
				if err != nil {
					return errors.New("file does not exist")
				}
				if !stat.Mode().IsRegular() {
					return errors.New("not a file")
				}
				return nil
			},
		}
		result, err := prompt.Run()
		if err != nil {
			logger.Info("Prompt failed: ", err)
		}
		errAbsPath := utils.AbsPath(&result)
		if errAbsPath != nil {
			logger.Info("Failed to get absolute path of SSH key: ", errAbsPath)
		}
		sshKeyPath = result
		logger.Debug("Using SSH key path: ", sshKeyPath)
	}

	if sshKeyPath != "" && sshKeyPassphrase == "" && !isSilent {
		prompt := promptui.Prompt{
			Label:     "SSH Key Passphrase",
			AllowEdit: true,
			Mask:      '*',
		}
		result, err := prompt.Run()
		if err != nil {
			logger.Info("Prompt failed: ", err)
		}
		sshKeyPassphrase = result
	}

	if sshKeyPath != "" {
		logger.Debug("Trying SSH NewPublicKeysFromFile with user provided key")
		am, errAuth := ssh.NewPublicKeysFromFile(u, sshKeyPath, sshKeyPassphrase)
		if errAuth != nil {
			logger.Fatal("Failed to create SSH agent: ", errAuth)
		}
		refsDefaultAuth, errDefaultAuth := remote.List(&gogit.ListOptions{
			Auth: am,
		})
		if errDefaultAuth != nil {
			logger.Fatal("Failed to get branches from remote: ", errDefaultAuth)
		}
		authMethod = am
		logger.Info("Using user provided SSH key")
		logger.Info("Successfully authenticated with SSH key")
		remoteRefs = refsDefaultAuth
	}
}

func tryNewSSHAgentAuth(u string) bool {
	logger.Debug("Trying SSH NewSSHAgentAuth")
	defaultAuth, errDefaultAuthBuilder := ssh.DefaultAuthBuilder(u)
	if errDefaultAuthBuilder != nil {
		logger.Info("Failed to create SSH agent auth: ", errDefaultAuthBuilder)
		return false
	}
	refsDefaultAuth, errDefaultAuth := remote.List(&gogit.ListOptions{
		Auth: defaultAuth,
	})
	if errDefaultAuth != nil {
		logger.Info("Unable to use default auth method", errDefaultAuth)
		return false
	}
	authMethod = defaultAuth
	remoteRefs = refsDefaultAuth
	logger.Info("Using default SSH agent auth")
	logger.Info("Successfully authenticated with SSH agent")
	return true
}

func tryHostNameKeyConfig(h string, u string) bool {
	logger.Debug("Trying SSH NewPublicKeysFromFile from default user settings")
	identityFile := ssh.DefaultSSHConfig.Get(h, "IdentityFile")
	errAbs := utils.AbsPath(&identityFile)
	if errAbs != nil {
		logger.Info("Failed to get absolute path of identity file: ", errAbs)
		return false
	}
	logger.Debug("Using identity file: ", identityFile)
	am, errAuth := ssh.NewPublicKeysFromFile(u, identityFile, "")
	if errAuth != nil {
		logger.Info("Failed to create SSH agent auth: ", errAuth)
		return false
	}
	refsDefaultAuth, errDefaultAuth := remote.List(&gogit.ListOptions{
		Auth: am,
	})
	if errDefaultAuth != nil {
		logger.Info("Failed to get branches from remote: ", errDefaultAuth)
		return false
	}
	authMethod = am
	remoteRefs = refsDefaultAuth
	logger.Info("Using default SSH keys")
	logger.Info("Successfully authenticated with default SSH key")
	return true
}
