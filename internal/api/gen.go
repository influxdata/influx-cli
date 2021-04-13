package api

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -package api -generate types -o types.gen.go api.yml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -package api -generate client -o client.gen.go api.yml
