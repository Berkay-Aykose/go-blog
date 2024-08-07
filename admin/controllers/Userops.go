package controllers

import (
	"crypto/sha256"
	"fmt"
	"goblog/admin/helpers"
	"goblog/admin/models"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

type Userops struct{}

func (userops Userops) Index(c *gin.Context) {
	view, err := template.ParseFiles(helpers.Include("userops/login")...)

	if err != nil {
		println(err)
		return
	}

	data := make(map[string]interface{})
	data["Alert"] = helpers.GetAlert(c)
	view.ExecuteTemplate(c.Writer, "index", data)
}

func (userops Userops) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(c.PostForm("password"))))

	user := models.User{}.Get("username = ? AND password = ?", username, password)

	if (user.Username == username) && (user.Password == password) {
		helpers.SetUser(c, username, password)
		helpers.SetAlert(c, "Hoşgeldiniz...")
		c.Redirect(http.StatusSeeOther, "/admin")
	} else {
		helpers.SetAlert(c, "yanlış Kullanıcı Adı veya Şifre...")
		c.Redirect(http.StatusSeeOther, "/admin/login")
	}
}

func (userops Userops) Logout(c *gin.Context) {
	helpers.RemoveUser(c)
	helpers.SetAlert(c, "Hoşçakalın...")
	c.Redirect(http.StatusSeeOther, "/admin/login")

}
