package middleware

import (
	"errors"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	icontext "github.com/influxdata/influx-cli/v2/pkg/cli/context"
	"github.com/urfave/cli"
)

const (
	OSSBuildHeader   = "OSS"
	CloudBuildHeader = "Cloud"
)

func WrongHostErrString(commandHost, actualHost string) string {
	return fmt.Sprintf("Error: InfluxDB %s-only command used with InfluxDB %s host", commandHost, actualHost)
}

var HandleExit cli.ExitErrHandlerFunc = func(ctx *cli.Context, err error) {
	if err == nil {
		return
	}

	var header string
	var genericErr api.GenericOpenAPIError
	if errors.As(err, &genericErr) {
		header = genericErr.BuildHeader()
	}

	// Replace the error message with the relevant information if a platform-specific command was used on the wrong host.
	// Otherwise, pass the error message along as-is to the CLI exit handler.
	var setErr bool
	if header == OSSBuildHeader {
		if icontext.GetCloudOnly(ctx) {
			err = cli.NewExitError(WrongHostErrString(CloudBuildHeader, OSSBuildHeader), 1)
			setErr = true
		}
	} else if header == CloudBuildHeader {
		if icontext.GetOssOnly(ctx) {
			err = cli.NewExitError(WrongHostErrString(OSSBuildHeader, CloudBuildHeader), 1)
			setErr = true
		}
	}
	if !setErr {
		err = cli.NewExitError(fmt.Sprintf("Error: %v", err.Error()), 1)
	}

	cli.HandleExitCoder(err)
}
