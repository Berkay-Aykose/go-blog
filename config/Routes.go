package config

import (
	"goblog/admin/controllers"
	site "goblog/site/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	// Serve static files
	r.Static("/admin/assets", "./admin/assets")
	r.Static("/site/assets", "./site/assets")
	r.Static("/uploads", "./uploads")

	// Admin routes
	adminRoutes := r.Group("/admin")
	{
		admin := controllers.Dashboard{}
		userops := controllers.Userops{}
		categories := controllers.Categories{}

		// Admin dashboard
		adminRoutes.GET("/", admin.Index)
		// Admin add new item page
		adminRoutes.GET("/yeni-ekle", admin.NewItem)
		// Add new item form
		adminRoutes.POST("/add", admin.Add)
		// Edit blog
		adminRoutes.GET("/edit/:id", admin.Edit)
		// Delete blog
		adminRoutes.GET("/delete/:id", admin.Delete)
		// Update blog
		adminRoutes.POST("/update/:id", admin.Update)

		// Admin login
		adminRoutes.GET("/login", userops.Index)
		// Admin login form
		adminRoutes.POST("/do-login", userops.Login)
		// Admin logout
		adminRoutes.GET("/logout", userops.Logout)

		// Admin categories
		adminRoutes.GET("/kategoriler/", categories.Index)
		adminRoutes.POST("/kategoriler/add", categories.Add)
		adminRoutes.GET("/kategoriler/delete/:id", categories.Delete)
	}

	// Site routes
	siteRoutes := r.Group("/")
	{
		homepage := site.Homepage{}

		// Homepage
		siteRoutes.GET("/", homepage.Index)
		// Blog details
		siteRoutes.GET("/yazilar/:slug", homepage.Detail)
	}

	return r
}
