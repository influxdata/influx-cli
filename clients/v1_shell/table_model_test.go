package v1shell

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/stretchr/testify/assert"
)

func Test_Explore(t *testing.T) {
	r, w, e := os.Pipe()
	if e != nil {
		assert.FailNow(t, e.Error())
	}
	old := os.Stdout
	os.Stdout = w
	defer func() {
		os.Stdout = old
	}()

	print()
	seriesName := "tagless"
	seriesTags := map[string]string{"state": ""}
	seriesPartial := false
	seriesColumns := []string{"time", "state", "value"}
	seriesValues := [][]interface{}{{1731069761000000000, "on", 2},
		{1731069771000000000, "off", 1}, {1731069781000000000, "on", 0}}

	model := NewModel(
		api.InfluxqlJsonResponseSeries{
			Name:    &seriesName,
			Tags:    &seriesTags,
			Partial: &seriesPartial,
			Columns: &seriesColumns,
			Values:  &seriesValues,
		},
		true,
		"test",
		map[string]string{"foo": ""},
		0,
		10,
		0,
		10,
		false,
	)

	buffer := &bytes.Buffer{}
	go func() {
		_, err := io.Copy(buffer, r)
		if err != nil {
			assert.FailNow(t, err.Error())
		}
	}()
	model.Init()
	model.View()
	// Tried WG and Chan but both seem to get hung-up with tea for some reason
	time.Sleep(100 * time.Millisecond)
	os.Stdout = old
	check := buffer.String()
	checkLines := strings.Split(check, "\n")
	assert.Equal(t, "Name: test", checkLines[0])
	assert.Equal(t, "Tags: foo=N/A", checkLines[1])
}
