package web

import (
	"net/http"

	"github.com/7wik-pk/BnB-bookingsapp/pkg/config"
	"github.com/7wik-pk/BnB-bookingsapp/pkg/handlers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func setupRouter(app *config.AppConfig) (router *gin.Engine) {
	router = gin.Default()

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

	return router
}
