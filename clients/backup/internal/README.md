# V1 Meta Protobufs

For compatibility with backups made via the v2.0.x `influx` CLI, we include logic
for opening & reading backed-up KV stores to derive bucket manifests. Part of that
process requires reading & unmarshalling V1 database info, serialized as protobuf.
To support that requirement, we've copied the `meta.proto` definition out of `influxdb`
and into this repository. This file isn't intended to be modified.

If `meta.pb.go` ever needs to be re-generated, follow these steps:
1. Install `protoc` (i.e. via `brew install protobuf`)
2. Run `go install github.com/gogo/protobuf/protoc-gen-gogo` from within this repository
3. Run `go generate <path to clients/backup>`
