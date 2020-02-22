package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func (app *AppContainer) GraphQLHandler(h http.Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := c.Request()
		response := c.Response()

		h.ServeHTTP(response, request)
		return nil
	}
}
