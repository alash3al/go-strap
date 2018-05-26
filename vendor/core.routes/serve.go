package routes

import (
	"core.globals"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

// CustomValidator used in in Echo as main validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate  Validates the specified struct
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// Serve the routes
func Serve() {
	// Echo instance
	e := echo.New()

	// Some settings
	e.HideBanner = true
	e.HidePort = true
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// register the system routes
	e.GET("/", Hello)

	// Start server
	globals.LogSuccess("⇨ Starting HTTP server (%s) \n", globals.ENV["SERVER_HTTP_ADDR"])
	globals.LogErr("⇨ [Server] %s", e.Start(globals.ENV["SERVER_HTTP_ADDR"]).Error())
}
