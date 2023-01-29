package blueprint

import (
	"github.com/labstack/echo/v4"
)

type SuccessResponse struct {
	Response bool                   `json:"response"`
	Message  string                 `json:"message"`
	Data     map[string]interface{} `json:"data"`
}

type Controller interface {
	Create(c echo.Context) *SuccessResponse
}
