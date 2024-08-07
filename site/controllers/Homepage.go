package controllers

import (
	"fmt"
	"goblog/site/helpers"
	"goblog/site/models"
	"net/http"
	"text/template"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Homepage struct{}

func (homepage Homepage) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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
	view.ExecuteTemplate(w, "index", data)
}

func (homepage Homepage) Detail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("homepage/detail")...)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	slug := params.ByName("slug")
	data["Post"] = models.Post{}.Get("slug = ?", slug)
	view.ExecuteTemplate(w, "detail", data)
}
