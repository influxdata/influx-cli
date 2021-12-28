# influx-cli

CLI for managing resources in InfluxDB v2

## Motivation

This repository decouples the `influx` CLI from the OSS `influxdb` codebase. Our goals are to:
1. Make it easier to keep the CLI up-to-date with InfluxDB Cloud API changes
2. Enable faster turn-around on fixes/features that only affect the CLI
3. Allow the CLI to be built & released for a wider range of platforms than the server can support

## Updating dependencies

We use [`OpenAPITools/openapi-generator`](https://github.com/OpenAPITools/openapi-generator) to generate
the underlying HTTP client used by the CLI.
`influx-cli/openapi` is a Git submodule. If you change or update your branch, you may need to update and regenerate _openapi_.
You'll need Docker running locally for the generate script to work.

To update, run the following commands in your `influx-cli` repo:

1. Update the _openapi_ Git submodule.

   `git pull --recurse-submodules`
   
2. With Docker running locally, regenerate _openapi_.

   `make openapi`

## Building the CLI

1. Clone this repo (influx-cli) and change to your _influx-cli_ directory.

   ```
   git clone git@github.com:influxdata/influx-cli.git
   cd influx-cli
   ```
   
2. If you're building a different version or commit, you may need to update and rebuild the _openapi_ git submodule.

   ```
   git pull --recurse-submodules
   make openapi
   ```
   
3. Build the CLI.
   
   `make` or `make influx`

The `make` command writes the binary to `bin/$(GOOS)/influx`.

## Running the CLI

After building, use `influx -h` to see the list of available commands.

### Enabling Completions

The CLI supports generating completions for `bash`, `zsh`, and `powershell`. To enable completions for a
single shell session, run one of these commands:
```
# For bash:
source <(influx completion bash)
# For zsh:
source <(influx completion zsh)
# For pwsh:
Invoke-Expression ((influx completion powershell) -join "`n`")
```
To enable completions across sessions, add the appropriate line to your shell's login profile (i.e. `~/.bash_profile`).

## Testing

Run `make test` to run unit tests.
