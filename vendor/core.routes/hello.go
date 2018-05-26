package routes

import (
	"core.globals"
	"github.com/labstack/echo"
)

// Hello send a welcome message to the requester
func Hello(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"success": true,
		"message": "Welcome to " + globals.ENV["SYSTEM_NAME"],
	})
}
