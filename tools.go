//+build tools

package influxcli

// This package is a workaround for adding additional paths to the go.mod file
// and ensuring they stay there. The build tag ensures this file never gets
// compiled, but the go module tool will still look at the dependencies and
// add/keep them in go.mod so we can version these paths along with our other
// dependencies. When we run build on any of these paths, we get the version
// that has been specified in go.mod rather than the master copy.

import (
	_ "github.com/daixiang0/gci"
	_ "github.com/golang/mock/mockgen"
	_ "golang.org/x/tools/cmd/goimports"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
