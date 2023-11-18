//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=types.cfg.yaml ../openapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yaml ../openapi.yaml

package api
