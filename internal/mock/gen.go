package mock

//go:generate mockgen -package mock -destination api_bucket_schemas.gen.go github.com/influxdata/influx-cli/v2/internal/api BucketSchemasApi
//go:generate mockgen -package mock -destination api_buckets.gen.go github.com/influxdata/influx-cli/v2/internal/api BucketsApi
