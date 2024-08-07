package helpers

import (
	"goblog/admin/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func SetUser(c *gin.Context, username string, password string) error {
	session, err := store.Get(c.Request, "blog-user")
	if err != nil {
		println(err)
		return err
	}
	session.Values["username"] = username
	session.Values["password"] = password

	return sessions.Save(c.Request, c.Writer)
}

func CheckUser(c *gin.Context) bool {
	session, err := store.Get(c.Request, "blog-user")
	if err != nil {
		println(err)
		return false
	}
	username := session.Values["username"]
	password := session.Values["password"]

	user := models.User{}.Get("username = ? AND password = ?", username, password)

	if (user.Username == username) && (user.Password == password) {
		return true
	}
	return false
}

func RemoveUser(c *gin.Context) error {
	session, err := store.Get(c.Request, "blog-user")
	if err != nil {
		println(err)
		return err
	}

	session.Options.MaxAge = -1

	return sessions.Save(c.Request, c.Writer)
}
