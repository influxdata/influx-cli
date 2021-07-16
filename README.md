# influx-cli

CLI for managing resources in InfluxDB v2

## Motivation

This repository decouples the `influx` CLI from the OSS `influxdb` codebase. Our goals are to:
1. Make it easier to keep the CLI up-to-date with InfluxDB Cloud API changes
2. Enable faster turn-around on fixes/features that only affect the CLI
3. Allow the CLI to be built & released for a wider range of platforms than the server can support

## Building

Run `make` or `make influx` to build the CLI. The output binary will be written to `bin/$(GOOS)/influx`.

### Regenerating OpenAPI client

We use [`OpenAPITools/openapi-generator`](https://github.com/OpenAPITools/openapi-generator) to generate
the underlying HTTP client used by the CLI. Run `make openapi` to re-generate the code. You'll need Docker
running locally for the script to work.

## Running

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
