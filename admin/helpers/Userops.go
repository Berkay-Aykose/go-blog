package helpers

import (
	"goblog/admin/models"
	"net/http"

	"github.com/gorilla/sessions"
)

func SetUser(w http.ResponseWriter, r *http.Request, username string, password string) error {
	session, err := store.Get(r, "blog-user")
	if err != nil {
		println(err)
		return err
	}
	session.Values["username"] = username
	session.Values["password"] = password

	return sessions.Save(r, w)
}

func CheckUser(w http.ResponseWriter, r *http.Request) bool {
	session, err := store.Get(r, "blog-user")
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

func RemoveUser(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "blog-user")
	if err != nil {
		println(err)
		return err
	}

	session.Options.MaxAge = -1

	return sessions.Save(r, w)
}
