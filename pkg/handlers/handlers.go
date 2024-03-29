package handlers

import (
	"log"
	"net/http"

	"github.com/7wik-pk/wayne-manor-bnb/pkg/config"
	"github.com/7wik-pk/wayne-manor-bnb/pkg/models"
	"github.com/7wik-pk/wayne-manor-bnb/templates"
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

		session.Set(remoteIPKey, remoteIP)
		if err := session.Save(); err != nil {
			log.Println("error encountered while saving user IP address to session: ", err.Error())
			ctx.String(http.StatusInternalServerError, "500: Internal Server Error")
			return
		}

		ctx.HTML(http.StatusOK, templates.HomePage, &models.TemplateData{
			Title: homePageTitle,
			StringMap: map[string]string{
				welcomeMessage: "Welcome",
			},
		})
	} else {
		// if session exists i.e, user is revisiting the page in an existing session

		ctx.HTML(http.StatusOK, templates.HomePage, &models.TemplateData{
			Title: homePageTitle,
			StringMap: map[string]string{
				welcomeMessage: "Welcome Back",
			},
		})
	}

}

// Handler for the about page
func (repo *repository) About(ctx *gin.Context) {

	// log.Println("inside handlers.About()")
	ctx.HTML(http.StatusOK, templates.AboutPage, &models.TemplateData{Title: aboutPageTitle})

}

// Handler for the page that provides available rooms based on input parameters.
func (repo *repository) AvailableRooms(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, templates.AvailableRoomsPage, &models.TemplateData{Title: availableRoomsPageTitle})
}

// Handler for the contact page
func (repo *repository) Contact(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, templates.ContactPage, &models.TemplateData{Title: contactPageTitle})
}

// Handler for non-existing pages
func (repo *repository) NoRoute(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, templates.NotFoundPage, &models.TemplateData{Title: notFoundPageTitle})
}
