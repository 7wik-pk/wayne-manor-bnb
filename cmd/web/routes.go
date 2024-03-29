package web

import (
	"net/http"

	"github.com/7wik-pk/wayne-manor-bnb/pkg/config"
	"github.com/7wik-pk/wayne-manor-bnb/pkg/handlers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	csrf "github.com/utrack/gin-csrf"
)

func setupRouter(app *config.AppConfig) (router *gin.Engine) {
	router = gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/**/*.gohtml")

	// configuring middlewares
	router.Use(sessions.Sessions("session", app.CookieStore))
	router.Use(csrf.Middleware(csrf.Options{
		Secret: app.CsrfSecret,
		ErrorFunc: func(ctx *gin.Context) {
			ctx.String(http.StatusBadRequest, "CSRF token mismatch")
		},
	}))

	// configuring routes
	router.GET("/", handlers.Repo.Home)

	router.GET("/about", handlers.Repo.About)
	router.GET("/contact", handlers.Repo.Contact)

	router.GET("/rooms/available", handlers.Repo.AvailableRooms)
	// router.GET("/rooms/cancel", )

	router.NoRoute(handlers.Repo.NoRoute)

	return router
}
