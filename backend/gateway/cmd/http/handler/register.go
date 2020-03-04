package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/purwandi/hazelapp/user/types"
)

// RegisterHandler is handle user registration
func (app *AppContainer) RegisterHandler(c echo.Context) error {
	input := types.RegisterUserInput{
		Fullname:             c.FormValue("fullname"),
		Email:                c.FormValue("email"),
		Username:             c.FormValue("username"),
		Password:             c.FormValue("password"),
		PasswordConfirmation: c.FormValue("password_confirmation"),
	}

	user, err := app.Resolver.UserService.RegisterUser(input)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, user)
}
