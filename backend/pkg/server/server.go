package server

import (
	"backend/pkg/router"
	"fmt"

	"github.com/labstack/echo/v4"
)

func ConnectServer() {
	e := echo.New()
	router.RegisterRouter(e)
	e.Logger.Fatal(e.Start(":3223"))

	fmt.Println("Server Start")
}
