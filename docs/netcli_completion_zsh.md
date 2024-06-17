## netcli completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(netcli completion zsh)

To load completions for every new session, execute once:

#### Linux:

	netcli completion zsh > "${fpath[1]}/_netcli"

#### macOS:

	netcli completion zsh > $(brew --prefix)/share/zsh/site-functions/_netcli

You will need to start a new shell for this setup to take effect.


```
netcli completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --debug-logging   Enable debug logging. This can be set using the environment variable DEBUG=true.
  -s, --silent          Silent mode. Do not prompt for any input.
```

### SEE ALSO

* [netcli completion](netcli_completion.md)	 - Generate the autocompletion script for the specified shell

