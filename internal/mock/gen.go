package mock

// HTTP API mocks
//go:generate go run github.com/golang/mock/mockgen -package mock -destination api_bucket_schemas.gen.go github.com/influxdata/influx-cli/v2/internal/api BucketSchemasApi
//go:generate go run github.com/golang/mock/mockgen -package mock -destination api_buckets.gen.go github.com/influxdata/influx-cli/v2/internal/api BucketsApi
//go:generate go run github.com/golang/mock/mockgen -package mock -destination api_health.gen.go github.com/influxdata/influx-cli/v2/internal/api HealthApi
//go:generate go run github.com/golang/mock/mockgen -package mock -destination api_organizations.gen.go github.com/influxdata/influx-cli/v2/internal/api OrganizationsApi
//go:generate go run github.com/golang/mock/mockgen -package mock -destination api_setup.gen.go github.com/influxdata/influx-cli/v2/internal/api SetupApi
//go:generate go run github.com/golang/mock/mockgen -package mock -destination api_write.gen.go github.com/influxdata/influx-cli/v2/internal/api WriteApi

// Other mocks
//go:generate go run github.com/golang/mock/mockgen -package mock -destination config.gen.go -mock_names Service=MockConfigService github.com/influxdata/influx-cli/v2/internal/config Service
//go:generate go run github.com/golang/mock/mockgen -package mock -destination stdio.gen.go github.com/influxdata/influx-cli/v2/internal/stdio StdIO
