package mock

// HTTP API mocks
//go:generate go run github.com/golang/mock/mockgen -package mock -destination api_bucket_schemas.gen.go github.com/influxdata/influx-cli/v2/api BucketSchemasApi
//go:generate go run github.com/golang/mock/mockgen -package mock -destination api_buckets.gen.go github.com/influxdata/influx-cli/v2/api BucketsApi
//go:generate go run github.com/golang/mock/mockgen -package mock -destination api_health.gen.go github.com/influxdata/influx-cli/v2/api HealthApi
//go:generate go run github.com/golang/mock/mockgen -package mock -destination api_organizations.gen.go github.com/influxdata/influx-cli/v2/api OrganizationsApi
//go:generate go run github.com/golang/mock/mockgen -package mock -destination api_setup.gen.go github.com/influxdata/influx-cli/v2/api SetupApi
//go:generate go run github.com/golang/mock/mockgen -package mock -destination api_write.gen.go github.com/influxdata/influx-cli/v2/api WriteApi
//go:generate go run github.com/golang/mock/mockgen -package mock -destination api_query.gen.go github.com/influxdata/influx-cli/v2/api QueryApi
//go:generate go run github.com/golang/mock/mockgen -package mock -destination api_users.gen.go github.com/influxdata/influx-cli/v2/api UsersApi
//go:generate go run github.com/golang/mock/mockgen -package mock -destination api_delete.gen.go github.com/influxdata/influx-cli/v2/api DeleteApi

// Other mocks
//go:generate go run github.com/golang/mock/mockgen -package mock -destination config.gen.go -mock_names Service=MockConfigService github.com/influxdata/influx-cli/v2/internal/config Service
//go:generate go run github.com/golang/mock/mockgen -package mock -destination stdio.gen.go github.com/influxdata/influx-cli/v2/pkg/stdio StdIO
