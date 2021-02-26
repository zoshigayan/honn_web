package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type BookController struct {
}

func (b BookController) Init(g *echo.Group) {
	g.GET("", b.Index)
}

func (b BookController) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "book/index", nil)
}
