package controllers

import (
	"goblog/admin/helpers"
	"goblog/admin/models"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type Categories struct {
}

// anasayfa admin
func (categories Categories) Index(c *gin.Context) {
	if !helpers.CheckUser(c) {
		helpers.SetAlert(c, "Lütfen giriş yapınız...")
		c.Redirect(http.StatusSeeOther, "/admin/login")
		return
	}
	view, err := template.ParseFiles(helpers.Include("categories/list")...)
	if err != nil {
		println(err)
		return
	}

	data := make(map[string]interface{})
	data["Categories"] = models.Category{}.GetAll()
	data["Alert"] = helpers.GetAlert(c)

	view.ExecuteTemplate(c.Writer, "index", data)
}

func (categories Categories) Add(c *gin.Context) {
	if !helpers.CheckUser(c) {
		helpers.SetAlert(c, "Lütfen giriş yapınız...")
		c.Redirect(http.StatusSeeOther, "/admin/login")
		return
	}

	categoryTitle := c.PostForm("category-title")
	categorySlug := slug.Make(categoryTitle)

	models.Category{
		Title: categoryTitle,
		Slug:  categorySlug,
	}.Add()

	helpers.SetAlert(c, "Kayıt Başarıyla Eklendi...")
	c.Redirect(http.StatusSeeOther, "/admin/kategoriler")
}

func (categories Categories) Delete(c *gin.Context) {
	if !helpers.CheckUser(c) {
		helpers.SetAlert(c, "Lütfen giriş yapınız...")
		c.Redirect(http.StatusSeeOther, "/admin/login")
		return
	}

	category := models.Category{}.Get(c.Params.ByName("id"))
	category.Delete()
	//Kayıt olunca eski sayfaya dönemk için
	c.Redirect(http.StatusSeeOther, "/admin/kategoriler")
}
