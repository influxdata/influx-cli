package clients

import (
	"fmt"

	"github.com/influxdata/influx-cli/v2/pkg/influxid"
)

type OrgParams struct {
	OrgID   influxid.ID
	OrgName string
}

type BucketParams struct {
	BucketID   influxid.ID
	BucketName string
}

type OrgBucketParams struct {
	OrgParams
	BucketParams
}

type AuthLookupParams struct {
	ID       influxid.ID
	Username string
}

func (p AuthLookupParams) Validate() (err error) {
	if p.Username == "" && !p.ID.Valid() {
		err = fmt.Errorf("id or username required")
	} else if p.Username != "" && p.ID.Valid() {
		err = fmt.Errorf("specify id or username, not both")
	}
	return
}

func (p AuthLookupParams) IsSet() bool {
	return p.ID.Valid() || p.Username != ""
}
