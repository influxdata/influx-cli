package mock

import (
	"context"

	"github.com/influxdata/influx-cli/v2/internal/api"
)

var _ api.BucketsApi = (*BucketsApi)(nil)

type BucketsApi struct {
	DeleteBucketsIDExecuteFn func(api.ApiDeleteBucketsIDRequest) error
	GetBucketsExecuteFn      func(api.ApiGetBucketsRequest) (api.Buckets, error)
	GetBucketsIDExecuteFn    func(api.ApiGetBucketsIDRequest) (api.Bucket, error)
	PatchBucketsIDExecuteFn  func(api.ApiPatchBucketsIDRequest) (api.Bucket, error)
	PostBucketsExecuteFn     func(api.ApiPostBucketsRequest) (api.Bucket, error)
}

func (b *BucketsApi) DeleteBucketsID(_ context.Context, bucketID string) api.ApiDeleteBucketsIDRequest {
	return api.ApiDeleteBucketsIDRequest{ApiService: b}.BucketID(bucketID)
}

func (b *BucketsApi) DeleteBucketsIDExecute(r api.ApiDeleteBucketsIDRequest) error {
	return b.DeleteBucketsIDExecuteFn(r)
}

func (b *BucketsApi) GetBuckets(context.Context) api.ApiGetBucketsRequest {
	return api.ApiGetBucketsRequest{ApiService: b}
}

func (b *BucketsApi) GetBucketsExecute(r api.ApiGetBucketsRequest) (api.Buckets, error) {
	return b.GetBucketsExecuteFn(r)
}

func (b *BucketsApi) GetBucketsID(_ context.Context, bucketID string) api.ApiGetBucketsIDRequest {
	return api.ApiGetBucketsIDRequest{ApiService: b}.BucketID(bucketID)
}

func (b *BucketsApi) GetBucketsIDExecute(r api.ApiGetBucketsIDRequest) (api.Bucket, error) {
	return b.GetBucketsIDExecuteFn(r)
}

func (b *BucketsApi) PatchBucketsID(_ context.Context, bucketID string) api.ApiPatchBucketsIDRequest {
	return api.ApiPatchBucketsIDRequest{ApiService: b}.BucketID(bucketID)
}

func (b *BucketsApi) PatchBucketsIDExecute(r api.ApiPatchBucketsIDRequest) (api.Bucket, error) {
	return b.PatchBucketsIDExecuteFn(r)
}

func (b *BucketsApi) PostBuckets(context.Context) api.ApiPostBucketsRequest {
	return api.ApiPostBucketsRequest{ApiService: b}
}

func (b *BucketsApi) PostBucketsExecute(r api.ApiPostBucketsRequest) (api.Bucket, error) {
	return b.PostBucketsExecuteFn(r)
}
