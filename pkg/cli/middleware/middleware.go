package middleware

import (
	"github.com/urfave/cli/v2"
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
