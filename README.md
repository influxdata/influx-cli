# influx-cli

CLI for managing resources in InfluxDB v2

## Status

This is a work-in-progress effort to decouple the `influx` CLI from the OSS `influxdb` codebase.
Our goals are to:
1. Make it easier to keep the CLI up-to-date with InfluxDB Cloud API changes
2. Enable faster turn-around on fixes/features that only affect the CLI
3. Allow the CLI to be built & released for a wider range of platforms than the server can support

## Building

Run `make` to build the CLI. The output binary will be written to `bin/$(GOOS)/influx`.

### Regenerating OpenAPI client

We use [`OpenAPITools/openapi-generator`](https://github.com/OpenAPITools/openapi-generator) to generate
the underlying HTTP client used by the CLI. Run `make openapi` to re-generate the code. You'll  need Docker
running locally for the script to work.

## Running

After building, use `influx -h` to see the list of available commands.

### Enabling Completions

The CLI supports generating completions for `bash`, `zsh`, and `powershell`:
```
# For bash:
source <(influx completion bash)
# For zsh:
source <(influx completion zsh)
# For pwsh:
Invoke-Expression ((influx completion powershell) -join "`n`")
```

## Testing

Run `make test` to run unit tests.
