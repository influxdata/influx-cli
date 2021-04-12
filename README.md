# influx-cli

CLI for managing resources in InfluxDB v2

## Status

This is a work-in-progress effort to decouple the `influx` CLI from the OSS `influxdb` codebase.
Our goals are to:
1. Make it easier to keep the CLI up-to-date with InfluxDB Cloud API changes
2. Enable faster turn-around on fixes/features that only affect the CLI
3. Allow the CLI to be built & released for a wider range of platforms than the server can support
