package influxid

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"strconv"
)

// IDLength is the exact length a string (or a byte slice representing it) must have in order to be decoded into a valid ID.
const IDLength = 16

var (
	// ErrInvalidID signifies invalid IDs.
	ErrInvalidID = errors.New("invalid ID")

	// ErrInvalidIDLength is returned when an ID has the incorrect number of bytes.
	ErrInvalidIDLength = errors.New("id must have a length of 16 bytes")
)

// ID is a unique identifier.
//
// Its zero value is not a valid ID.
type ID uint64

// IDFromString creates an ID from a given string.
//
// It errors if the input string does not match a valid ID.
func IDFromString(str string) (ID, error) {
	var id ID
	err := id.DecodeFromString(str)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// MustIDFromString is like IDFromString but panics if
// s is not a valid base-16 identifier.
func MustIDFromString(s string) ID {
	id, err := IDFromString(s)
	if err != nil {
		panic(err)
	}
	return id
}

// Decode parses b as a hex-encoded byte-slice-string.
//
// It errors if the input byte slice does not have the correct length
// or if it contains all zeros.
func (i *ID) Decode(b []byte) error {
	if len(b) != IDLength {
		return ErrInvalidIDLength
	}

	res, err := strconv.ParseUint(string(b), 16, 64)
	if err != nil {
		return ErrInvalidID
	}

	if *i = ID(res); !i.Valid() {
		return ErrInvalidID
	}
	return nil
}

// DecodeFromString parses s as a hex-encoded string.
func (i *ID) DecodeFromString(s string) error {
	return i.Decode([]byte(s))
}

// Encode converts ID to a hex-encoded byte-slice-string.
//
// It errors if the receiving ID holds its zero value.
func (i ID) Encode() ([]byte, error) {
	if !i.Valid() {
		return nil, ErrInvalidID
	}

	b := make([]byte, hex.DecodedLen(IDLength))
	binary.BigEndian.PutUint64(b, uint64(i))

	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b)
	return dst, nil
}

// Valid checks whether the receiving ID is a valid one or not.
func (i ID) Valid() bool {
	return i != 0
}

// String returns the ID as a hex encoded string.
//
// Returns an empty string in the case the ID is invalid.
func (i ID) String() string {
	enc, _ := i.Encode()
	return string(enc)
}

// GoString formats the ID the same as the String method.
// Without this, when using the %#v verb, an ID would be printed as a uint64,
// so you would see e.g. 0x2def021097c6000 instead of 02def021097c6000
// (note the leading 0x, which means the former doesn't show up in searches for the latter).
func (i ID) GoString() string {
	return `"` + i.String() + `"`
}

// MarshalText encodes i as text.
// Providing this method is a fallback for json.Marshal,
// with the added benefit that IDs encoded as map keys will be the expected string encoding,
// rather than the effective fmt.Sprintf("%d", i) that json.Marshal uses by default for integer types.
func (i ID) MarshalText() ([]byte, error) {
	return i.Encode()
}

// UnmarshalText decodes i from a byte slice.
// Providing this method is also a fallback for json.Unmarshal,
// also relevant when IDs are used as map keys.
func (i *ID) UnmarshalText(b []byte) error {
	return i.Decode(b)
}

func (i *ID) Set(s string) error {
	id, err := IDFromString(s)
	if err != nil {
		return err
	}
	*i = id
	return nil
}
