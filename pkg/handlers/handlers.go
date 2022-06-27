package handlers

import (
	"log"
	"net/http"

	"github.com/7wik-pk/BnB-bookingsapp/pkg/config"
	"github.com/7wik-pk/BnB-bookingsapp/pkg/models"
	"github.com/7wik-pk/BnB-bookingsapp/pkg/render"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type repository struct {
	App *config.AppConfig
}

var Repo *repository

func NewRepository(app *config.AppConfig) *repository {
	return &repository{
		App: app,
	}
}

func Init(repo *repository) {
	Repo = repo
}

// Handler for the Home page
func (repo *repository) Home(ctx *gin.Context) {
	remoteIP := ctx.Request.RemoteAddr

	session := sessions.Default(ctx)

	if session.Get("remote_ip") == nil {
		session.Set("remote_ip", remoteIP)
		if err := session.Save(); err != nil {
			log.Println("error encountered while saving user IP address to session: ", err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "internal server error",
			})
			return
		}
	}

	render.Template(ctx, "home.page.tmpl", &models.TemplateData{})

}

// Handler for the about page
// func (repo *repository) About(ctx *gin.Context) {}
