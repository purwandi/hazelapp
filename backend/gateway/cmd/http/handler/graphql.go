package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
)

func (app *AppContainer) GraphQLHandler(h http.Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := c.Request()
		response := c.Response()

		auth := request.Header.Get("Authorization")
		bearer := len("Bearer")

		if len(auth) < bearer+1 || auth[:bearer] != "Bearer" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "missing or malformed authrozation",
			})
		}

		token := auth[bearer+1:]
		result, err := app.Resolver.UserService.FindUserByAccessToken(token)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "This endpoint requires you to be authenticated.",
			})
		}

		ctx := context.WithValue(request.Context(), "viewer", result)

		h.ServeHTTP(response, request.WithContext(ctx))
		return nil
	}
}
