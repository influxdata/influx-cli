## v2.5.0 [2022-10-28]
----------------------

### Bug Fixes

1. [3a593e7](https://github.com/influxdata/influx-cli/commit/3a593e7): Don't allow creating an auth with instance resources
1. [e2aa4d2](https://github.com/influxdata/influx-cli/commit/e2aa4d2): Fix stack error typo
1. [77cca94](https://github.com/influxdata/influx-cli/commit/77cca94): Fixes an error where stdin could not be used to create tasks
1. [121864a](https://github.com/influxdata/influx-cli/commit/121864a): Cloud expects dataType in csv files
1. [3285a03](https://github.com/influxdata/influx-cli/commit/3285a03): Do not require remoteOrgID for remote connection creation

### Features

1. [6142b7a](https://github.com/influxdata/influx-cli/commit/6142b7a): Support username-password in config update

### Other

1. [e39e365](https://github.com/influxdata/influx-cli/commit/e39e365): Chore: go1.19
1. [188c393](https://github.com/influxdata/influx-cli/commit/188c393): Chore: fix go version in `go.mod`


## v2.4.0 [2022-08-18]
----------------------

### Bug Fixes

1. [50de408](https://github.com/influxdata/influx-cli/commit/50de408): Add mutual exclusion for OrgId and OrgName params
1. [0c17ebd](https://github.com/influxdata/influx-cli/commit/0c17ebd): Users and orgs permissions should not be scoped under an org
1. [d3e0efb](https://github.com/influxdata/influx-cli/commit/d3e0efb):
1. [b9ffcb4](https://github.com/influxdata/influx-cli/commit/b9ffcb4): Improve display for strings and numbers in v1 shell tables
1. [182303e](https://github.com/influxdata/influx-cli/commit/182303e): Prevent v1 shell hang on empty query result
1. [75e3606](https://github.com/influxdata/influx-cli/commit/75e3606): Show `remotes` and `replications` flags in `auth create`
1. [fbbe974](https://github.com/influxdata/influx-cli/commit/fbbe974): Update unsupported xcode version

### Features

1. [30e64c5](https://github.com/influxdata/influx-cli/commit/30e64c5): Add --extra-http-header flag
1. [fc52974](https://github.com/influxdata/influx-cli/commit/fc52974): Add back the InfluxQL REPL
1. [760f07e](https://github.com/influxdata/influx-cli/commit/760f07e): Invokable scripts
1. [9dc1b8e](https://github.com/influxdata/influx-cli/commit/9dc1b8e): Add pretty table format to REPL
1. [e5707cd](https://github.com/influxdata/influx-cli/commit/e5707cd): Allow setting membership type in influx org members add
1. [da2899d](https://github.com/influxdata/influx-cli/commit/da2899d): Add skipRowOnError handling for raw line protocol files
1. [d470527](https://github.com/influxdata/influx-cli/commit/d470527): Added tag stripping step to openapi generation to fix codegen
1. [0de05ed](https://github.com/influxdata/influx-cli/commit/0de05ed): Updated openapi to support tasks containing scripts
1. [f34e6a8](https://github.com/influxdata/influx-cli/commit/f34e6a8): Add username and password login
1. [0b6ce21](https://github.com/influxdata/influx-cli/commit/0b6ce21): Allow deleting replications with remotes
1. [1453e20](https://github.com/influxdata/influx-cli/commit/1453e20): Added script support when creating tasks for the cloud
1. [826e03f](https://github.com/influxdata/influx-cli/commit/826e03f): Added script support when updating tasks for the cloud
1. [7bdad28](https://github.com/influxdata/influx-cli/commit/7bdad28): Add virtual column to DBRP printing
1. [5c7c34f](https://github.com/influxdata/influx-cli/commit/5c7c34f): Replication bucket name

### Other

1. [3527951](https://github.com/influxdata/influx-cli/commit/3527951): Build: upgrade to Go 1.18.1
1. [51ca97e](https://github.com/influxdata/influx-cli/commit/51ca97e): Build: upgrade to Go 1.18.3
1. [c695e60](https://github.com/influxdata/influx-cli/commit/c695e60): Add REPL autocompletion & go-prompt
1. [a68106e](https://github.com/influxdata/influx-cli/commit/a68106e): Replace token flags with = to prevent bad parsing of leading dash in token
1. [09881c0](https://github.com/influxdata/influx-cli/commit/09881c0): Chore: fix typo in mockgen
1. [85c690f](https://github.com/influxdata/influx-cli/commit/85c690f): Chore: add checkgenerate test to `lint`
1. [56a8276](https://github.com/influxdata/influx-cli/commit/56a8276): Chore: fix issues with Go 1.18 in CI
1. [c44d2f2](https://github.com/influxdata/influx-cli/commit/c44d2f2): Build: upload "latest" artifacts
1. [78ef3c1](https://github.com/influxdata/influx-cli/commit/78ef3c1): Chore: add influx command shell hint
1. [65ff49f](https://github.com/influxdata/influx-cli/commit/65ff49f): Chore(tasks): support looking up tasks by script id
1. [051a6aa](https://github.com/influxdata/influx-cli/commit/051a6aa): Clarify difference in virtual vs physical dbrps when listing


## v2.3.0 [2022-03-18]
----------------------

### Bug Fixes

1. [6a7c4f6](https://github.com/influxdata/influx-cli/commit/6a7c4f6): `v1 auth create` supports multiple buckets
1. [becbe8f](https://github.com/influxdata/influx-cli/commit/becbe8f): Use influx-debug-id header
1. [13d0827](https://github.com/influxdata/influx-cli/commit/13d0827): Duration parser shows duration missing units on error
1. [9ddf110](https://github.com/influxdata/influx-cli/commit/9ddf110): Template apply uses better diff checking
1. [37ec38a](https://github.com/influxdata/influx-cli/commit/37ec38a): Rename bucket id parameters to be explicit
1. [c8c7c1c](https://github.com/influxdata/influx-cli/commit/c8c7c1c): Json suffix for json template from CLI

### Features

1. [81de7e6](https://github.com/influxdata/influx-cli/commit/81de7e6): Return error if API token required but not found
1. [99791ba](https://github.com/influxdata/influx-cli/commit/99791ba): Add flags for remotes, replications, and functions to `auth create`
1. [f32a55f](https://github.com/influxdata/influx-cli/commit/f32a55f): Add `drop-non-retryable-data` to replications commands
1. [4c0fae3](https://github.com/influxdata/influx-cli/commit/4c0fae3): Add ExecuteWithHttpInfo methods for generated API
1. [327f239](https://github.com/influxdata/influx-cli/commit/327f239): Enable remotes and replication streams feature
1. [178c754](https://github.com/influxdata/influx-cli/commit/178c754): Add server-config command
1. [7af0b2a](https://github.com/influxdata/influx-cli/commit/7af0b2a): Enhanced error messages for cloud and oss specific commands
1. [88ba346](https://github.com/influxdata/influx-cli/commit/88ba346): Add max age to replications create and update

### Other

1. [a3af8ca](https://github.com/influxdata/influx-cli/commit/a3af8ca): Revert: "feat: return error if API token required but not found
1. [566dcaf](https://github.com/influxdata/influx-cli/commit/566dcaf): Chore: update CHANGELOG
1. [1eadcf1](https://github.com/influxdata/influx-cli/commit/1eadcf1): Chore: update CHANGELOG
1. [adc58b8](https://github.com/influxdata/influx-cli/commit/adc58b8): Chore: refactor `influxid.ID`, cleanup organization checking
1. [a408c02](https://github.com/influxdata/influx-cli/commit/a408c02): Chore: remove remote validation
1. [68ac116](https://github.com/influxdata/influx-cli/commit/68ac116): Chore: update openapi ref to latest
1. [cb3bade](https://github.com/influxdata/influx-cli/commit/cb3bade): Update: instructions to update openapi
1. [5cd1c9d](https://github.com/influxdata/influx-cli/commit/5cd1c9d): Build: automatically generate changelog
1. [981be78](https://github.com/influxdata/influx-cli/commit/981be78): Chore: upgrade to latest OpenAPI
1. [041ebf6](https://github.com/influxdata/influx-cli/commit/041ebf6): Chore: improve logging for creating a new remote
1. [85a33ad](https://github.com/influxdata/influx-cli/commit/85a33ad): Chore: update to API spec for new template endpoint response


## v2.2.1 [2021-11-09]
----------------------

### Bug Fixes

1. [0364b18](https://github.com/influxdata/influx-cli/commit/0364b18): Return "unknown command" instead of "no help topic" error when unknown (sub)command is passed
1. [74a1b8e](https://github.com/influxdata/influx-cli/commit/74a1b8e): Wrong position of orgID and userID in `org members remove`

### Other

1. [31ac783](https://github.com/influxdata/influx-cli/commit/31ac783): Chore: pin date for 2.2.1 release

## v2.2.0 [2021-10-21]
----------------------

### Features

1. [259](https://github.com/influxdata/influx-cli/pull/259): Add `-b` shorthand for `--bucket` to `delete`
1. [285](https://github.com/influxdata/influx-cli/pull/285): Add short-hand `--all-access` and `--operator` options to `auth create`.
1. [307](https://github.com/influxdata/influx-cli/pull/307): Handle pagination in `bucket list` to support displaying more than 20 buckets.

### Bug Fixes

1. [297](https://github.com/influxdata/influx-cli/pull/297): Detect and warn when `restore --full` changes the operator token.
1. [302](https://github.com/influxdata/influx-cli/pull/302): Set newly-created config as active in `setup`.
1. [305](https://github.com/influxdata/influx-cli/pull/305): Embed timezone data into Windows builds to avoid errors.

## v2.1.1 [2021-09-22]
----------------------

### Go Version

This release upgrades the project to `go` version 1.17.

### Bug Fixes

1. [221](https://github.com/influxdata/influx-cli/pull/221): Fix shell completion for top-level `influx` commands.
1. [228](https://github.com/influxdata/influx-cli/pull/228): Make global `--http-debug` flag visible in help text.
1. [232](https://github.com/influxdata/influx-cli/pull/232): Don't set empty strings for IDs in permission resources.
1. [236](https://github.com/influxdata/influx-cli/pull/236): Detect and error out on incorrect positional args.
1. [255](https://github.com/influxdata/influx-cli/pull/255): Respect value of `--host` flag when writing CLI configs in `setup`.
1. [272](https://github.com/influxdata/influx-cli/pull/272): Reduce overuse of "token" in help text.

## v2.1.0 [2021-07-29]
----------------------

### New Repository

This is the initial release of the `influx` CLI from the [`influxdata/influx-cli`](https://github.com/influxdata/influx-cli/)
GitHub repository.

### Breaking Changes

#### `influx write` skip-header parsing

The `write` command no longer supports `--skipHeader` as short-hand for `--skipHeader 1`. This change was made to
simplify our CLI parser.

#### Stricter input validation for template-related commands

The `apply`, `export`, and `stacks` commands now raise errors when CLI options fail to parse  instead of silently
discarding bad inputs. This change was made to help users debug when their commands fail to execute as expected.

#### Server-side template summarization & validation

The `template` and `template validate` commands now use an API request to the server to perform their logic,
instead of performing the work on the client-side. Offline summarization & validation is no longer supported.
This change was made to avoid significant code duplication between `influxdb` and `influx-cli`, and to allow server-
side template logic to evolve without requiring coordinated CLI changes.

#### `influx stacks --json` output conventions

The output of `influx stacks --json` previously used an UpperCamelCase naming convention for most, but not all, keys.
The command now uses lowerCamelCase consistently for all objects keys, matching the schema returned by the API.

### Features

1. [33](https://github.com/influxdata/influx-cli/pull/33): Add global `--http-debug` flag to help inspect communication with InfluxDB servers.
1. [52](https://github.com/influxdata/influx-cli/pull/52): Add `bucket-schema` commands to manage explicit measurement schemas in InfluxDB Cloud.
1. [52](https://github.com/influxdata/influx-cli/pull/52): Update `bucket create` to allow setting a schema type.
1. [52](https://github.com/influxdata/influx-cli/pull/52): Update `bucket list` to display schema types.
1. [116](https://github.com/influxdata/influx-cli/pull/116): Reimplement `backup` to support downloading embedded SQL store from InfluxDB v2.1.x.
1. [116](https://github.com/influxdata/influx-cli/pull/116): Add `--compression` flag to `backup` to support enabling/disabling GZIP compression of downloaded files.
1. [121](https://github.com/influxdata/influx-cli/pull/121): Reimplement `restore` to support uploading embedded SQL store from InfluxDB v2.1.x.
1. [191](https://github.com/influxdata/influx-cli/pull/191): Add `--password` flag to `user password` command to allow bypassing interactive prompt.
1. [208](https://github.com/influxdata/influx-cli/pull/208): Bind `--skip-verify` flag to `INFLUX_SKIP_VERIFY` environment variable.

### Bug Fixes

1. [35](https://github.com/influxdata/influx-cli/pull/35): Fix interactive password collection & color rendering in PowerShell.
1. [97](https://github.com/influxdata/influx-cli/pull/97): `org members list` no longer hangs on organizations with more than 10 members.
1. [109](https://github.com/influxdata/influx-cli/pull/109): Detect & warn when inputs to `write` contain standalone CR characters.
1. [122](https://github.com/influxdata/influx-cli/pull/122): `dashboards` command now accepts `--org` flag, or falls back to default org in config.
1. [140](https://github.com/influxdata/influx-cli/pull/140): Return a consistent error when responses fail to decode, with hints for OSS- our Cloud-only commands.
