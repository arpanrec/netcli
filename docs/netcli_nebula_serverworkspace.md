## netcli nebula serverworkspace

Setup workspace for development using server workspace playbook

### Synopsis

Setup workspace for development using

[server workspace playbook](https://github.com/arpanrec/arpanrec.nebula/blob/main/playbooks/server_workspace.md)

```
netcli nebula serverworkspace [flags]
```

### Options

```
      --bitwarden-desktop    Install Bitwarden Desktop
      --bws                  Install BWS
      --code                 Install Visual Studio Code
      --go                   Install GoLang
  -h, --help                 help for serverworkspace
      --java                 Install Java
      --mattermost-desktop   Install Mattermost Desktop
      --nodejs               Install Node.js
      --postman              Install Postman
      --pulumi               Install Pulumi
      --raw string           Pass raw arguments to the script. Example: --raw "--nodejs --go --java", this will also add the local config file: .tmp/serverworkspace-local-config.json
      --telegram-desktop     Install Telegram Desktop
      --terminal             Install Terminal
      --terraform            Install Terraform
      --themes               Install Themes
      --vault                Install Vault
```

### Options inherited from parent commands

```
      --debug-logging   Enable debug logging. This can be set using the environment variable DEBUG=true.
  -s, --silent          Silent mode. Do not prompt for any input.
```

### SEE ALSO

* [netcli nebula](netcli_nebula.md)	 - Nebula Runner

