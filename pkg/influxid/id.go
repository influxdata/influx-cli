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

// Validate ensures that a passed string has a valid ID syntax.
// Checks that the string is of length 16, and is a valid hex-encoded uint.
func Validate(id string) error {
	_, err := Decode(id)
	return err
}

// Encode converts a uint64 to a hex-encoded byte-slice-string.
func Encode(id uint64) string {
	b := make([]byte, hex.DecodedLen(IDLength))
	binary.BigEndian.PutUint64(b, id)

	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b)
	return string(dst)
}

// Decode parses id as a hex-encoded byte-slice-string.
//
// It errors if the input byte slice does not have the correct length
// or if it contains all zeros.
func Decode(id string) (uint64, error) {
	if len([]byte(id)) != 16 {
		return 0, ErrInvalidIDLength
	}
	res, err := strconv.ParseUint(id, 16, 64)
	if err != nil || res == 0 {
		return 0, ErrInvalidID
	}
	return res, nil
}
