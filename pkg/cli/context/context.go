package context

import "github.com/urfave/cli"

const (
	contextKeyCloudOnly = "cloudOnly"
	contextKeyOssOnly   = "ossOnly"
)

func SetCloudOnly(ctx *cli.Context) {
	ctx.App.Metadata[contextKeyCloudOnly] = true
}

func SetOssOnly(ctx *cli.Context) {
	ctx.App.Metadata[contextKeyOssOnly] = true
}

func GetCloudOnly(ctx *cli.Context) bool {
	_, ok := ctx.App.Metadata[contextKeyCloudOnly].(bool)
	return ok
}

func GetOssOnly(ctx *cli.Context) bool {
	_, ok := ctx.App.Metadata[contextKeyOssOnly].(bool)
	return ok
}
