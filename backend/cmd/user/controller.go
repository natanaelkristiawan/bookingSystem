package user

import (
	"backend/internal/blueprint"

	"github.com/labstack/echo/v4"
)

type User struct {
	firstName string
	lastName  string
	email     string
}

// post url
func (param User) Create(c echo.Context) *blueprint.SuccessResponse {

	var person = map[string]interface{}{
		"firstName": c.FormValue("firstName"),
		"lastName":  c.FormValue("lastName"),
		"email":     c.FormValue("email"),
	}

	respone := blueprint.SuccessResponse{
		Response: true,
		Message:  "Success create data",
		Data:     person,
	}

	return &respone
}
