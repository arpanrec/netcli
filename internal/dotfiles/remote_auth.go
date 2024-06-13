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
)

func createRemoteAuth() {
	remote = gogit.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		URLs: []string{RepositoryUrl},
	})

	gitURL, errUrlParse := giturl.Parse(RepositoryUrl)
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

	if schema == "https" {
		logger.Debug("Using HTTPS auth method")
		refsDefaultAuth, errDefaultAuth := getRefs(nil)
		if errDefaultAuth != nil {
			logger.Fatal("Failed to get branches from remote: ", errDefaultAuth)
		}
		authMethod = nil
		remoteRefs = refsDefaultAuth
		return
	}

	logger.Fatal("Unsupported schema: ", schema)
}

func tryWithUserProvidedKey(u *string) bool {

	if SshKeyPath == "" && !isSilent {
		prompt := promptui.Prompt{
			Label:     "SSH Key Path (optional)",
			AllowEdit: true,
			Validate: func(s string) error {
				if s == "" {
					return nil
				}
				return utils.ValidateFile(s, false)
			},
		}
		result, err := prompt.Run()
		if err != nil {
			utils.IsInterrupt(&err)
			logger.Info("Prompt failed: ", err)
		}
		if result != "" {
			absPath, errAbsPath := utils.AbsPath(result)
			if errAbsPath != nil {
				logger.Info("Failed to get absolute path of SSH key: ", errAbsPath)
			}
			SshKeyPath = absPath
			logger.Debug("Using SSH key path: ", SshKeyPath)
		} else {
			logger.Info("No SSH key path provided")
			return false
		}
	}

	if SshKeyPath == "" {
		return false
	}

	if SshKeyPath != "" && SshKeyPassphrase == "" && !isSilent && !sshKeyPassphraseProvided {
		prompt := promptui.Prompt{
			Label:     "SSH Key Passphrase for " + SshKeyPath + " (optional)",
			AllowEdit: true,
			Mask:      '*',
		}
		result, err := prompt.Run()
		if err != nil {
			utils.IsInterrupt(&err)
			logger.Info("Prompt failed: ", err)
		}
		SshKeyPassphrase = result
	}

	logger.Debug("Trying SSH with user provided key: ", SshKeyPath)
	am, errAuth := ssh.NewPublicKeysFromFile(*u, SshKeyPath, SshKeyPassphrase)
	if errAuth != nil {
		logger.Fatal("Failed to create SSH agent: ", errAuth)
	}
	refsDefaultAuth, errDefaultAuth := getRefs(am)
	if errDefaultAuth != nil {
		logger.Fatal("Failed to get branches from remote: ", errDefaultAuth)
	}
	authMethod = am
	remoteRefs = refsDefaultAuth
	logger.Info("Successfully authenticated with user provided SSH key: ", SshKeyPath)
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

	abs, errAbs := utils.AbsPath(identityFile)
	if errAbs != nil {
		logger.Warn("Failed to get absolute path of identity file: ", errAbs)
		return false
	}
	identityFile = abs

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
