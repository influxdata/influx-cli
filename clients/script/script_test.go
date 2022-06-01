package script_test

import (
	"context"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/script"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_SimpleCreate(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	scriptsApi := mock.NewMockInvocableScriptsApi(ctrl)

	var (
		scriptId       = "123456789"
		scriptName     = "simple"
		scriptDesc     = "A basic script to be created"
		scriptOrgId    = "1111111111111"
		scriptContent  = `from(bucket: "sample_data") |> range(start: -10h)`
		scriptLanguage = api.SCRIPTLANGUAGE_FLUX
	)

	scriptsApi.EXPECT().PostScripts(gomock.Any()).Return(api.ApiPostScriptsRequest{
		ApiService: scriptsApi,
	})

	scriptsApi.EXPECT().PostScriptsExecute(gomock.Any()).Return(api.Script{
		Id:          &scriptId,
		Name:        scriptName,
		Description: &scriptDesc,
		OrgID:       scriptOrgId,
		Script:      scriptContent,
		Language:    &scriptLanguage,
	}, nil)

	stdio := mock.NewMockStdIO(ctrl)
	client := script.Client{
		CLI:                 clients.CLI{StdIO: stdio, PrintAsJSON: true},
		InvocableScriptsApi: scriptsApi,
	}

	stdio.EXPECT().Write(tmock.MatchedBy(func(in []byte) bool {
		t.Logf("Stdio output: %s", in)
		inStr := string(in)
		// Verify we print the basic details of the script in some form.
		return strings.Contains(inStr, scriptId) &&
			strings.Contains(inStr, scriptName) &&
			strings.Contains(inStr, scriptOrgId)
	}))

	params := script.CreateParams{
		Description: scriptDesc,
		Language:    string(scriptLanguage),
		Name:        scriptName,
		Script:      scriptContent,
	}

	require.NoError(t, client.Create(context.Background(), &params))
}
