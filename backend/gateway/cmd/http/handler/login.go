package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/purwandi/hazelapp/user/types"
)

func (app *AppContainer) LoginHandler(c echo.Context) error {
	input := types.LoginUserInput{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}

	user, err := app.Resolver.UserService.LoginUser(input)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, user)
}
