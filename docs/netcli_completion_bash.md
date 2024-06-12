## netcli completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(netcli completion bash)

To load completions for every new session, execute once:

#### Linux:

	netcli completion bash > /etc/bash_completion.d/netcli

#### macOS:

	netcli completion bash > $(brew --prefix)/etc/bash_completion.d/netcli

You will need to start a new shell for this setup to take effect.


```
netcli completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --debug-logging   Enable debug logging. This can be set using the environment variable DEBUG=true.
  -s, --silent          Silent mode. Do not prompt for any input.
```

### SEE ALSO

* [netcli completion](netcli_completion.md)	 - Generate the autocompletion script for the specified shell

