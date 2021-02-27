package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/zoshigayan/honn_web/db"
	"github.com/zoshigayan/honn_web/models"
	"net/http"
	"strconv"
	"time"
)

type BookController struct {
}

func (b BookController) Init(g *echo.Group) {
	g.GET("", b.Index)
	g.GET("/:id", b.Show)
	g.GET("/new", b.New)
	g.POST("/create", b.Create)
}

func (b BookController) Index(c echo.Context) error {
	db := db.DbManager()

	books := []models.Book{}
	db.Find(&books)

	return c.Render(http.StatusOK, "book/index", map[string]interface{}{
		"Books": books,
	})
}

func (b BookController) Show(c echo.Context) error {
	db := db.DbManager()

	book := models.Book{}
	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&book, id)
	return c.Render(http.StatusOK, "book/show", map[string]interface{}{
		"Book": book,
	})
}

func (b BookController) New(c echo.Context) error {
	return c.Render(http.StatusOK, "book/new", nil)
}

func (b BookController) Create(c echo.Context) error {
	db := db.DbManager()

	authors := []models.Author{
		models.Author{Name: c.FormValue("Author")},
	}

	publishedAt, err := time.Parse("2006-01-02", c.FormValue("PublishedAt"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "error")
	}

	book := models.Book{
		Thumbnail:   c.FormValue("Thumbnail"),
		Title:       c.FormValue("Title"),
		Subtitle:    c.FormValue("Subtitle"),
		Authors:     authors,
		PublishedAt: publishedAt,
	}
	db.Create(&book)

	return c.Redirect(http.StatusSeeOther, "/books")
}
