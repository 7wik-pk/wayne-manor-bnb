package config

import (
	"github.com/gin-contrib/sessions/cookie"
)

type AppConfig struct {
	InProduction bool
	PortNumber   string
	// TemplateCache map[string]*template.Template
	SessionKey  string
	CookieStore cookie.Store
	CsrfSecret  string
}
