package middlewares

import (
	"github.com/labstack/echo"
)

// Example middleware
func Example(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}
