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
	// log.Println("inside handlers.Home()")
	remoteIP := ctx.Request.RemoteAddr

	session := sessions.Default(ctx)

	if session.Get(remoteIPKey) == nil {
		// new session i.e, user is visiting the web app for the first time

		// log.Println("remote IP not set")
		session.Set(remoteIPKey, remoteIP)
		if err := session.Save(); err != nil {
			log.Println("error encountered while saving user IP address to session: ", err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				errorMessage: "Internal Server Error",
			})
			return
		}

		err := render.Template(ctx, "home.page.tmpl", &models.TemplateData{
			StringMap: map[string]string{
				templateMessage: "Hello There!",
			},
		})

		if err != nil {
			log.Println("error encountered while rendering template: ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				errorMessage: "Internal Server Error",
			})

			return
		}
	} else {
		// if session exists i.e, user is revisiting the page in an existing session
		err := render.Template(ctx, "home.page.tmpl", &models.TemplateData{
			StringMap: map[string]string{
				templateMessage: "Welcome Back!",
			},
		})

		if err != nil {
			log.Println("error encountered while rendering template: ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				errorMessage: "Internal Server Error",
			})

			return
		}
	}

}

// Handler for the about page
func (repo *repository) About(ctx *gin.Context) {
	remoteIP := ctx.Request.RemoteAddr
	session := sessions.Default(ctx)

	if session.Get(remoteIPKey) == nil {
		session.Set(remoteIPKey, remoteIP)
		if err := session.Save(); err != nil {
			log.Println("error encountered while saving user IP address to session: ", err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				errorMessage: "Internal Server Error",
			})
			return
		}

	}

	// log.Println("calling render.Template()")
	err := render.Template(ctx, "about.page.tmpl", &models.TemplateData{})
	if err != nil {
		log.Println("error encountered while rendering template: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			errorMessage: "Internal Server Error",
		})

		return
	}

}
