package controllers

import (
	"fmt"
	"goblog/admin/helpers"
	"goblog/admin/models"
	"io"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
)

type Dashboard struct {
}

// anasayfa admin
func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		helpers.SetAlert(w, r, "Lütfen giriş yapınız...")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}

	view, err := template.New("index").Funcs(template.FuncMap{
		"getCategory": func(categoryID int) string {
			return models.Category{}.Get(categoryID).Title
		},
	}).ParseFiles(helpers.Include("dashboard/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	data["Posts"] = models.Post{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

// ekleme sayfa admim
func (dashboard Dashboard) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		helpers.SetAlert(w, r, "Lütfen giriş yapınız...")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}

	view, err := template.ParseFiles(helpers.Include("dashboard/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	data["Categories"] = models.Category{}.GetAll()
	view.ExecuteTemplate(w, "indexAdd", data)
}

func (dashboard Dashboard) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		helpers.SetAlert(w, r, "Lütfen giriş yapınız...")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}

	title := r.FormValue("blog-title")
	slug := slug.Make(title)
	description := r.FormValue("blog-description")
	content := r.FormValue("blog-content")
	categoryID, _ := strconv.Atoi(r.FormValue("blog-category"))

	//UPLOAD
	r.ParseMultipartForm(10 << 20)
	File, header, err := r.FormFile("blog-pictures")

	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(f, File)
	//UPLOAD end

	if err != nil {
		fmt.Println(err)
		return
	}

	models.Post{
		Title:       title,
		Slug:        slug,
		Content:     content,
		CategoryID:  categoryID,
		Description: description,
		Picture_url: "uploads/" + header.Filename,
	}.Add()
	helpers.SetAlert(w, r, "Kayıt Başarıyla Eklendi") //alert mesajı indexe gönderiyor
	//Kayıt olunca eski sayfaya dönemk için
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

// edit sayfa admin
func (dashboard Dashboard) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		helpers.SetAlert(w, r, "Lütfen giriş yapınız...")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}

	view, err := template.ParseFiles(helpers.Include("dashboard/edit")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Post"] = models.Post{}.Get(params.ByName("id"))
	data["Categories"] = models.Category{}.GetAll()
	view.ExecuteTemplate(w, "indexEdit", data)
}

func (dashboard Dashboard) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		helpers.SetAlert(w, r, "Lütfen giriş yapınız...")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}

	post := models.Post{}.Get(params.ByName("id"))
	post.Delete()
	//Kayıt olunca eski sayfaya dönemk için
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (dashboard Dashboard) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		helpers.SetAlert(w, r, "Lütfen giriş yapınız...")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}

	post := models.Post{}.Get(params.ByName("id"))

	title := r.FormValue("blog-title")
	slug := slug.Make(title)
	description := r.FormValue("blog-description")
	content := r.FormValue("blog-content")
	categoryID, _ := strconv.Atoi(r.FormValue("blog-category"))
	is_selected := r.FormValue("is_selected")

	var picture_url string

	if is_selected == "1" {
		//UPLOAD
		r.ParseMultipartForm(10 << 20)
		File, header, err := r.FormFile("blog-pictures")

		if err != nil {
			fmt.Println(err)
			return
		}
		f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = io.Copy(f, File)
		if err != nil {
			fmt.Println(err)
			return
		}
		//UPLOAD end
		picture_url = "uploads/" + header.Filename
		os.Remove(post.Picture_url)
	} else {
		picture_url = post.Picture_url
	}

	post.Updates(models.Post{
		Title:       title,
		Slug:        slug,
		Content:     content,
		CategoryID:  categoryID,
		Description: description,
		Picture_url: picture_url,
	})
	http.Redirect(w, r, "/admin/edit/"+params.ByName("id"), http.StatusSeeOther)
}
