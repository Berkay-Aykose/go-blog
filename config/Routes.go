package config

import (
	admin "goblog/admin/controllers"
	site "goblog/site/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	//ADMİN index
	r.GET("/admin", admin.Dashboard{}.Index)
	//ADMİN blog ekleme sayfası
	r.GET("/admin/yeni-ekle", admin.Dashboard{}.NewItem)
	//blog ekle form
	r.POST("/admin/add", admin.Dashboard{}.Add)
	//ADMİN blog düzenle
	r.GET("/admin/edit/:id", admin.Dashboard{}.Edit)
	//ADMİN blog sil
	r.GET("/admin/delete/:id", admin.Dashboard{}.Delete)
	//ADMİN güncelle
	r.POST("/admin/update/:id", admin.Dashboard{}.Update)

	//Userops indexLogin
	r.GET("/admin/login", admin.Userops{}.Index)
	//AdMİN FORM login
	r.POST("/admin/do-login", admin.Userops{}.Login)
	//AdMİN session Logout
	r.GET("/admin/logout", admin.Userops{}.Logout)

	//Admin categories index
	r.GET("/admin/kategoriler/", admin.Categories{}.Index)
	//Admin kategori ekleme
	r.POST("/admin/kategoriler/add", admin.Categories{}.Add)
	//Admin kategori silme
	r.GET("/admin/kategoriler/delete/:id", admin.Categories{}.Delete)

	//Anasayfa
	r.GET("/", site.Homepage{}.Index)
	//Anasayfa yazılara gitme details
	r.GET("/yazilar/:slug", site.Homepage{}.Detail)

	//css kaynakları kullanmak için
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets")) //admin/assets/*tüm dosyalar gelise admin/assets yönlendir
	r.ServeFiles("/site/assets/*filepath", http.Dir("site/assets"))   //site/assets/*tüm dosyalar gelise site/assets yönlendir
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	return r
}
