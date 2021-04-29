package middleware

import (
	"github.com/urfave/cli/v2"
)

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

func WithAfterFns(fns ...cli.AfterFunc) cli.AfterFunc {
	return func(ctx *cli.Context) error {
		for _, fn := range fns {
			if err := fn(ctx); err != nil {
				return err
			}
		}
		return nil
	}
}
