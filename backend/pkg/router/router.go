package router

import (
	"backend/cmd/user"
	"backend/internal/blueprint"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRouter(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	var userController blueprint.Controller
	userController = user.User{}
	generateRoute("user", e, userController)

}

func generateRoute(routeName string, e *echo.Echo, controller blueprint.Controller) {
	e.POST(routeName, func(c echo.Context) error {
		response := controller.Create(c)

		fmt.Println(response)

		return c.JSON(http.StatusOK, response)
	})
}
