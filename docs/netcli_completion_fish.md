## netcli completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	netcli completion fish | source

To load completions for every new session, execute once:

	netcli completion fish > ~/.config/fish/completions/netcli.fish

You will need to start a new shell for this setup to take effect.


```
netcli completion fish [flags]
```

### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --debug-logging   Enable debug logging. This can be set using the environment variable DEBUG=true.
  -s, --silent          Silent mode. Do not prompt for any input.
```

### SEE ALSO

* [netcli completion](netcli_completion.md)	 - Generate the autocompletion script for the specified shell

