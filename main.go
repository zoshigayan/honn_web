package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zoshigayan/honn_web/config"
	"github.com/zoshigayan/honn_web/controllers"
	"github.com/zoshigayan/honn_web/db"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "assets")

	db.Init()

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/books")
	})
	controllers.BookController{}.Init(e.Group("/books"))

	e.Renderer = config.Renderer()
	e.Debug = true

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
