package web

import (
	"log"
	"net/http"
	"os"

	"github.com/7wik-pk/wayne-manor-bnb/pkg/config"
	"github.com/7wik-pk/wayne-manor-bnb/pkg/handlers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var app config.AppConfig

func Run() {

	// setting up the config variables

	app.PortNumber = os.Getenv("PORT")
	// app.PortNumber = "8080"
	app.InProduction = false
	app.SessionKey = "secret"
	app.CsrfSecret = "CSRFsecret"

	if app.InProduction {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	app.CookieStore = cookie.NewStore([]byte(app.SessionKey))
	app.CookieStore.Options(sessions.Options{
		Path:     "/",
		MaxAge:   0,
		Secure:   app.InProduction,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	repo := handlers.NewRepository(&app)
	handlers.Init(repo)

	initLogFile("./log/webapp.log")
	defer closeLogFile()

	logWriter, err := getLogWriter(true, true)
	if err != nil {
		log.Println("could not initialise logger, error: ", err.Error())
	}

	log.SetOutput(logWriter)
	gin.DefaultWriter = logWriter

	router := setupRouter(&app)

	if err := router.Run(":" + app.PortNumber); err != nil {
		log.Fatal("error: gin router crashed: ", err.Error())
	}

}
