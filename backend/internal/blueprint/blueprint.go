package blueprint

import (
	"github.com/labstack/echo/v4"
)

type SuccessResponse struct {
	Response bool
	Message  string
	Data     map[string]interface{}
}

type Controller interface {
	Create(c echo.Context) *SuccessResponse
}
