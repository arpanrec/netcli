package dotfiles

import (
	"github.com/arpanrec/netcli/internal/logger"
	giturl "github.com/chainguard-dev/git-urls"
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/go-git/go-git/v5/storage/memory"
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

		if tryWithUserProvidedKey(user) {
			return
		}

		if tryNewSSHAgentAuth(user) {
			return
		}

		if tryHostNameKeyConfig(hostname, user) {
			return
		}
	}

	logger.Debug(gitURL.Scheme)
	logger.Debug(gitURL.User.String())
	logger.Debug(gitURL.Port())
}

func tryWithUserProvidedKey(u string) bool {
	if sshKeyPath != "" {
		logger.Debug("Trying SSH NewPublicKeysFromFile")
		am, errAuth := ssh.NewPublicKeysFromFile(u, sshKeyPath, sshKeyPassphrase)
		if errAuth != nil {
			logger.Fatal("Failed to create SSH agent auth: ", errAuth)
			return false
		}
		refsDefaultAuth, errDefaultAuth := remote.List(&gogit.ListOptions{
			Auth: am,
		})
		if errDefaultAuth != nil {
			logger.Fatal("Failed to get branches from remote: ", errDefaultAuth)
			return false
		}
		authMethod = am
		remoteRefs = refsDefaultAuth
		return true
	}
	return false
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
	return true
}

func tryHostNameKeyConfig(h string, u string) bool {
	logger.Debug("Trying SSH NewPublicKeysFromFile")
	defaultUserSettings := ssh.DefaultSSHConfig.Get(h, "IdentityFile")

	am, errAuth := ssh.NewPublicKeysFromFile(u, defaultUserSettings, "")
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
	return true
}
