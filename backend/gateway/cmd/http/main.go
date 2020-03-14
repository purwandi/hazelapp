package main

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/purwandi/hazelapp/gateway/cmd/http/handler"
)

func main() {
	app := handler.NewAppContainer()

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Static("./public"))

	e.POST("/graphql", app.GraphQLHandler(&relay.Handler{
		Schema: graphql.MustParseSchema(app.Schema, app.Resolver),
	}))
	e.POST("/register", app.RegisterHandler)
	e.POST("/login", app.LoginHandler)

	e.Logger.Info(e.Start(":3001"))
}
