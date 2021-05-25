package clients

import "errors"

var (
	ErrPasswordIsTooShort = errors.New("password is too short")
	ErrMustSpecifyOrg     = errors.New("must specify org ID or org name")
	ErrMustSpecifyBucket  = errors.New("must specify bucket ID or bucket name")
)
