package dotfiles

import (
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

var repository *gogit.Repository

var remote *gogit.Remote

var authMethod transport.AuthMethod

var remoteRefs []*plumbing.Reference

var workTreeDir string

var backupDirRoot string
