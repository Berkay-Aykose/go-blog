package controllers

import (
	"fmt"
	"goblog/site/helpers"
	"goblog/site/models"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

type Homepage struct{}

func (homepage Homepage) Index(c *gin.Context) {
	view, err := template.New("index").Funcs(template.FuncMap{
		"getCategory": func(categoryID int) string {
			return models.Category{}.Get(categoryID).Title
		}, "getDate": func(t time.Time) string {
			return fmt.Sprintf("%d.%s.%d", t.Day(), t.Month().String(), t.Year())
		},
	}).ParseFiles(helpers.Include("homepage/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	data["Post"] = models.Post{}.GetAll()
	view.ExecuteTemplate(c.Writer, "index", data)
}

func (homepage Homepage) Detail(c *gin.Context) {
	view, err := template.ParseFiles(helpers.Include("homepage/detail")...)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	slug := c.Params.ByName("slug")
	data["Post"] = models.Post{}.Get("slug = ?", slug)
	view.ExecuteTemplate(c.Writer, "detail", data)
}
