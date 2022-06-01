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

func ptrFactory[T any](arg T) *T {
	return &arg
}

func Test_SimpleList(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	scriptsApi := mock.NewMockInvocableScriptsApi(ctrl)

	scripts := []api.Script{{
		Id:          ptrFactory("123456789"),
		Name:        "simple",
		Description: ptrFactory("First script"),
		OrgID:       "1111111111111",
		Script:      `from(bucket: "sample_data") |> range(start: -10h)`,
		Language:    ptrFactory(api.SCRIPTLANGUAGE_FLUX),
	}, {
		Id:          ptrFactory("000000001"),
		Name:        "another",
		Description: ptrFactory("Second script"),
		OrgID:       "9111111111119",
		Script:      `from(bucket: "sample_data") |> range(start: -5h)`,
		Language:    ptrFactory(api.SCRIPTLANGUAGE_FLUX),
	},
	}

	scriptsApi.EXPECT().GetScripts(gomock.Any()).Return(api.ApiGetScriptsRequest{
		ApiService: scriptsApi,
	})

	scriptsApi.EXPECT().GetScriptsExecute(gomock.Any()).Return(api.Scripts{
		Scripts: &scripts,
	}, nil)

	stdio := mock.NewMockStdIO(ctrl)
	client := script.Client{
		CLI:                 clients.CLI{StdIO: stdio, PrintAsJSON: true},
		InvocableScriptsApi: scriptsApi,
	}

	stdio.EXPECT().Write(tmock.MatchedBy(func(in []byte) bool {
		t.Logf("Stdio output: %s", in)
		inStr := string(in)
		// Verify we print the basic details of all scripts in some form.
		success := true
		for _, script := range scripts {
			success = success && strings.Contains(inStr, *script.Id)
			success = success && strings.Contains(inStr, script.Name)
			success = success && strings.Contains(inStr, script.OrgID)
		}
		return success
	}))

	params := script.ListParams{
		Limit:  10,
		Offset: 0,
	}

	require.NoError(t, client.List(context.Background(), &params))
}
