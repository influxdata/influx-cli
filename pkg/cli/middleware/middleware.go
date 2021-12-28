package middleware

import (
	icontext "github.com/influxdata/influx-cli/v2/pkg/cli/context"
	"github.com/urfave/cli"
)

// WithBeforeFns returns a cli.BeforeFunc that calls each of the provided
// functions in order.
// NOTE: The first function to return an error will end execution and
// be returned as the error value of the composed function.
func WithBeforeFns(fns ...cli.BeforeFunc) cli.BeforeFunc {
	return func(ctx *cli.Context) error {
		for _, fn := range fns {
			if err := fn(ctx); err != nil {
				return err
			}
		}
		return nil
	}
}

// AddMWToCmds is used to append a middleware to a list of existing commands.
func AddMWToCmds(cmds []cli.Command, mw cli.BeforeFunc) []cli.Command {
	newCmds := make([]cli.Command, 0, len(cmds))

	for _, cmd := range cmds {
		cmd.Before = WithBeforeFns(cmd.Before, mw)
		newCmds = append(newCmds, cmd)
	}

	return newCmds
}

var CloudOnly cli.BeforeFunc = func(ctx *cli.Context) error {
	icontext.SetCloudOnly(ctx)
	return nil
}

var OSSOnly cli.BeforeFunc = func(ctx *cli.Context) error {
	icontext.SetOssOnly(ctx)
	return nil
}
