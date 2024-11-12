package v1shell

import (
	"bytes"
	"context"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/stretchr/testify/assert"
)

type holder struct {
	buffer *bytes.Buffer
	mu     sync.Mutex
}

func Test_checkEmptyTagValueRender(t *testing.T) {

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

	h := &holder{buffer: &bytes.Buffer{}}
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1000)
	defer cancel()
	go func(ctx context.Context) {
		h.mu.Lock()
		_, err := io.CopyN(h.buffer, r, 29)
		h.mu.Unlock()
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		return
	}(ctx)

	model.Init()

	select {
	case <-ctx.Done():
		os.Stdout = old
		h.mu.Lock()
		check := h.buffer.String()
		h.mu.Unlock()
		checkLines := strings.Split(check, "\n")
		assert.Equal(t, "Name: test", checkLines[0])
		assert.Equal(t, "Tags: foo= ----- ", checkLines[1])
	}
}
