package controllers

import (
	"goblog/admin/helpers"
	"goblog/admin/models"
	"net/http"
	"text/template"

	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
)

type Categories struct {
}

// anasayfa admin
func (categories Categories) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		helpers.SetAlert(w, r, "Lütfen giriş yapınız...")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}
	view, err := template.ParseFiles(helpers.Include("categories/list")...)
	if err != nil {
		println(err)
		return
	}

	data := make(map[string]interface{})
	data["Categories"] = models.Category{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)

	view.ExecuteTemplate(w, "index", data)
}

func (categories Categories) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		helpers.SetAlert(w, r, "Lütfen giriş yapınız...")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}

	categoryTitle := r.FormValue("category-title")
	categorySlug := slug.Make(categoryTitle)

	models.Category{
		Title: categoryTitle,
		Slug:  categorySlug,
	}.Add()

	helpers.SetAlert(w, r, "Kayıt Başarıyla Eklendi...")
	http.Redirect(w, r, "/admin/kategoriler", http.StatusSeeOther)
}

func (categories Categories) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		helpers.SetAlert(w, r, "Lütfen giriş yapınız...")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}

	category := models.Category{}.Get(params.ByName("id"))
	category.Delete()
	//Kayıt olunca eski sayfaya dönemk için
	http.Redirect(w, r, "/admin/kategoriler", http.StatusSeeOther)
}
