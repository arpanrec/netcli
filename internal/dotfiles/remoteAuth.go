package dotfiles

import (
	"errors"
	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
	giturl "github.com/chainguard-dev/git-urls"
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/manifoldco/promptui"
	"os"
)

func createRemoteAuth() {
	remote = gogit.NewRemote(memory.NewStorage(), &config.RemoteConfig{
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

		if tryWithUserProvidedKey(&user) {
			return
		}

		if tryNewSSHAgentAuth(&user) {
			return
		}

		if tryHostNameKeyConfig(&hostname, &user) {
			return
		}
	}
	logger.Fatal("Failed to authenticate with remote")

}

func tryWithUserProvidedKey(u *string) bool {

	if sshKeyPath == "" && !isSilent {
		prompt := promptui.Prompt{
			Label:     "SSH Key Path (optional)",
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
			utils.IsInterrupt(&err)
			logger.Info("Prompt failed: ", err)
		}
		if result != "" {
			errAbsPath := utils.AbsPath(&result)
			if errAbsPath != nil {
				logger.Info("Failed to get absolute path of SSH key: ", errAbsPath)
			}
			sshKeyPath = result
			logger.Debug("Using SSH key path: ", sshKeyPath)
		} else {
			logger.Info("No SSH key path provided")
			return false
		}
	}

	if sshKeyPath == "" {
		return false
	}

	if sshKeyPath != "" && sshKeyPassphrase == "" && !isSilent && !sshKeyPassphraseProvided {
		prompt := promptui.Prompt{
			Label:     "SSH Key Passphrase for " + sshKeyPath + " (optional)",
			AllowEdit: true,
			Mask:      '*',
		}
		result, err := prompt.Run()
		if err != nil {
			utils.IsInterrupt(&err)
			logger.Info("Prompt failed: ", err)
		}
		sshKeyPassphrase = result
	}

	logger.Debug("Trying SSH with user provided key: ", sshKeyPath)
	am, errAuth := ssh.NewPublicKeysFromFile(*u, sshKeyPath, sshKeyPassphrase)
	if errAuth != nil {
		logger.Fatal("Failed to create SSH agent: ", errAuth)
	}
	refsDefaultAuth, errDefaultAuth := getRefs(am)
	if errDefaultAuth != nil {
		logger.Fatal("Failed to get branches from remote: ", errDefaultAuth)
	}
	authMethod = am
	remoteRefs = refsDefaultAuth
	logger.Info("Successfully authenticated with user provided SSH key: ", sshKeyPath)
	return true
}

func tryNewSSHAgentAuth(u *string) bool {
	logger.Debug("Trying SSH SSH Agent Auth")
	defaultAuth, errDefaultAuthBuilder := ssh.DefaultAuthBuilder(*u)
	if errDefaultAuthBuilder != nil {
		logger.Warn("Failed to create SSH agent auth: ", errDefaultAuthBuilder)
		return false
	}
	refsDefaultAuth, errDefaultAuth := getRefs(defaultAuth)
	if errDefaultAuth != nil {
		logger.Warn("Unable to use default auth method", errDefaultAuth)
		return false
	}
	authMethod = defaultAuth
	remoteRefs = refsDefaultAuth
	logger.Info("Successfully authenticated with SSH agent")
	return true
}

func tryHostNameKeyConfig(h *string, u *string) bool {
	logger.Debug("Trying SSH auth with hostname key config")
	identityFile := ssh.DefaultSSHConfig.Get(*h, "IdentityFile")
	errAbs := utils.AbsPath(&identityFile)
	if errAbs != nil {
		logger.Warn("Failed to get absolute path of identity file: ", errAbs)
		return false
	}
	logger.Debug("Using identity file: ", identityFile)
	am, errAuth := ssh.NewPublicKeysFromFile(*u, identityFile, "")
	if errAuth != nil {
		logger.Warn("Failed to get identity file: ", errAuth)
		return false
	}
	refsDefaultAuth, errDefaultAuth := getRefs(am)

	if errDefaultAuth != nil {
		logger.Warn("Failed to get branches from remote: ", errDefaultAuth)
		return false
	}
	authMethod = am
	remoteRefs = refsDefaultAuth
	logger.Info("Successfully authenticated with default SSH key")
	return true
}

func getRefs(auth transport.AuthMethod) ([]*plumbing.Reference, error) {
	allRefs, err := remote.List(&gogit.ListOptions{
		Auth: auth,
	})
	if err != nil {
		return nil, errors.New("failed to get branches from remote: " + err.Error())
	}
	return allRefs, nil
}
