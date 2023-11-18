package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/leandersteiner/project-management-backend-go/api"
	"github.com/leandersteiner/project-management-backend-go/handlers"
	middleware "github.com/oapi-codegen/echo-middleware"
)

func main() {
	port := flag.String("port", "8080", "Port for HTTP server")
	flag.Parse()

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Printf("Error loading swagger spec:\n %s", err)
		os.Exit(0)
	}

	swagger.Servers = nil

	h := handlers.New()

	e := echo.New()
	e.Use(echomiddleware.Logger())
	e.Use(middleware.OapiRequestValidator(swagger))

	api.RegisterHandlers(e, h)

	e.Logger.Fatal(e.Start(net.JoinHostPort("0.0.0.0", *port)))
}
