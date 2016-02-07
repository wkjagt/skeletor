package main

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func registerClient(e *echo.Echo) {
	e.Get("/bundle.js", func(c *echo.Context) error {
		return c.File("../client/bundle.js", "", false)
	})

	e.Get("/", func(c *echo.Context) error {
		return c.File("../client/index.html", "", false)
	})
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	registerClient(e)

	// Start server
	e.Run(":1323")
}
