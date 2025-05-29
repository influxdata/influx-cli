package influxid_test

import (
	"bytes"
	"testing"

	"github.com/influxdata/influx-cli/v2/pkg/influxid"
)

func TestDecode(t *testing.T) {
	if _, err := influxid.Decode("020f755c3c082000"); err != nil {
		t.Errorf("decode error: %v", err)
	}
}

func TestEncode(t *testing.T) {
	res, _ := influxid.Decode("5ca1ab1eba5eba11")
	want := []byte{53, 99, 97, 49, 97, 98, 49, 101, 98, 97, 53, 101, 98, 97, 49, 49}
	got := []byte(influxid.Encode(res))
	if !bytes.Equal(want, got) {
		t.Errorf("encoding error")
	}
}

func TestDecodeFromAllZeros(t *testing.T) {
	if _, err := influxid.Decode(string(make([]byte, influxid.IDLength))); err == nil {
		t.Errorf("expecting all zeros ID to not be a valid ID")
	}
}

func TestDecodeFromShorterString(t *testing.T) {
	if _, err := influxid.Decode("020f75"); err == nil {
		t.Errorf("expecting shorter inputs to error")
	}
}

func TestDecodeFromLongerString(t *testing.T) {
	if _, err := influxid.Decode("020f755c3c082000aaa"); err == nil {
		t.Errorf("expecting shorter inputs to error")
	}
}

func TestDecodeFromEmptyString(t *testing.T) {
	if _, err := influxid.Decode(""); err == nil {
		t.Errorf("expecting empty inputs to error")
	}
}
