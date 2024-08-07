package controllers

import (
	"fmt"
	"goblog/admin/helpers"
	"goblog/admin/models"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type Dashboard struct {
}

// anasayfa admin
func (dashboard Dashboard) Index(c *gin.Context) {
	if !helpers.CheckUser(c) {
		helpers.SetAlert(c, "Lütfen giriş yapınız...")
		c.Redirect(http.StatusSeeOther, "/admin/login")
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
	data["Alert"] = helpers.GetAlert(c)
	view.ExecuteTemplate(c.Writer, "index", data)
}

// ekleme sayfa admim
func (dashboard Dashboard) NewItem(c *gin.Context) {
	if !helpers.CheckUser(c) {
		helpers.SetAlert(c, "Lütfen giriş yapınız...")
		c.Redirect(http.StatusSeeOther, "/admin/login")
		return
	}

	view, err := template.ParseFiles(helpers.Include("dashboard/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	data["Categories"] = models.Category{}.GetAll()
	view.ExecuteTemplate(c.Writer, "indexAdd", data)
}

func (dashboard Dashboard) Add(c *gin.Context) {
	if !helpers.CheckUser(c) {
		helpers.SetAlert(c, "Lütfen giriş yapınız...")
		c.Redirect(http.StatusSeeOther, "/admin/login")
		return
	}

	title := c.PostForm("blog-title")
	slug := slug.Make(title)
	description := c.PostForm("blog-description")
	content := c.PostForm("blog-content")
	categoryID, _ := strconv.Atoi(c.PostForm("blog-category"))

	//UPLOAD
	const maxMemory = 10 << 20 // 10 MiB

	if err := c.Request.ParseMultipartForm(maxMemory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse form"})
		return
	}

	file, header, err := c.Request.FormFile("blog-pictures")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file from form"})
		return
	}
	defer file.Close()

	uploadsDir := "uploads"
	if err := os.MkdirAll(uploadsDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create uploads directory"})
		return
	}

	// Create the destination file
	dst, err := os.Create(filepath.Join(uploadsDir, header.Filename))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the file"})
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the destination file
	if _, err := io.Copy(dst, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the file"})
		return
	}
	//UPLOAD end

	models.Post{
		Title:       title,
		Slug:        slug,
		Content:     content,
		CategoryID:  categoryID,
		Description: description,
		Picture_url: "uploads/" + header.Filename,
	}.Add()
	helpers.SetAlert(c, "Kayıt Başarıyla Eklendi") //alert mesajı indexe gönderiyor
	//Kayıt olunca eski sayfaya dönemk için
	c.Redirect(http.StatusSeeOther, "/admin")
}

// edit sayfa admin
func (dashboard Dashboard) Edit(c *gin.Context) {
	if !helpers.CheckUser(c) {
		helpers.SetAlert(c, "Lütfen giriş yapınız...")
		c.Redirect(http.StatusSeeOther, "/admin/login")

		return
	}

	view, err := template.ParseFiles(helpers.Include("dashboard/edit")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Post"] = models.Post{}.Get(c.Params.ByName("id"))
	data["Categories"] = models.Category{}.GetAll()
	view.ExecuteTemplate(c.Writer, "indexEdit", data)
}

func (dashboard Dashboard) Delete(c *gin.Context) {
	if !helpers.CheckUser(c) {
		helpers.SetAlert(c, "Lütfen giriş yapınız...")
		c.Redirect(http.StatusSeeOther, "/admin/login")

		return
	}

	post := models.Post{}.Get(c.Params.ByName("id"))
	post.Delete()
	//Kayıt olunca eski sayfaya dönemk için
	c.Redirect(http.StatusSeeOther, "/admin")
}

func (dashboard Dashboard) Update(c *gin.Context) {
	if !helpers.CheckUser(c) {
		helpers.SetAlert(c, "Lütfen giriş yapınız...")
		c.Redirect(http.StatusSeeOther, "/admin/login")

		return
	}

	post := models.Post{}.Get(c.Params.ByName("id"))

	title := c.PostForm("blog-title")
	slug := slug.Make(title)
	description := c.PostForm("blog-description")
	content := c.PostForm("blog-content")
	categoryID, _ := strconv.Atoi(c.PostForm("blog-category"))
	is_selected := c.PostForm("is_selected")

	var picture_url string

	if is_selected == "1" {
		//UPLOAD
		const maxMemory = 10 << 20 // 10 MiB

		if err := c.Request.ParseMultipartForm(maxMemory); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse form"})
			return
		}
		file, header, err := c.Request.FormFile("blog-pictures")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file from form"})
			return
		}
		defer file.Close()

		f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = io.Copy(f, file)
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
	c.Redirect(http.StatusSeeOther, "/admin/edit/"+c.Params.ByName("id"))
}
