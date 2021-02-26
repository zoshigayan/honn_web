package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type BookController struct {
}

func (b BookController) Init(g *echo.Group) {
	g.GET("", b.Index)
	g.GET("/:id", b.Show)
}

func (b BookController) Index(c echo.Context) error {
	books := []struct {
		Id          string
		Thumbnail   string
		Title       string
		Subtitle    string
		Authors     []string
		PublishDate time.Time
		Description string
	}{
		{
			"hoge",
			"https://placehold.it/128x128?text=hoge",
			"ほげ",
			"ふが",
			[]string{"太郎", "花子"},
			time.Now(),
			"名著だよ～",
		},
	}
	return c.Render(http.StatusOK, "book/index", map[string]interface{}{
		"Books": books,
	})
}

func (b BookController) Show(c echo.Context) error {
	book := struct {
		Id          string
		Thumbnail   string
		Title       string
		Subtitle    string
		Authors     []string
		PublishDate time.Time
		Description string
	}{
		"hoge",
		"https://placehold.it/128x128?text=hoge",
		"ほげ",
		"ふが",
		[]string{"太郎", "花子"},
		time.Now(),
		"名著だよ～",
	}
	return c.Render(http.StatusOK, "book/show", map[string]interface{}{
		"Book": book,
	})
}
