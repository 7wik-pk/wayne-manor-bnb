package render

import (
	"bytes"
	"errors"
	"html/template"
	"path/filepath"

	"github.com/7wik-pk/BnB-bookingsapp/pkg/config"
	"github.com/7wik-pk/BnB-bookingsapp/pkg/models"
	"github.com/gin-gonic/gin"
)

var app *config.AppConfig
var templateFuncMap = template.FuncMap{}

func Init(appConfig *config.AppConfig) {
	app = appConfig
}

func Template(ctx *gin.Context, tmpl string, tmplData *models.TemplateData) error {
	// var templateCache map[string]*template.Template

	if !app.InProduction {
		if err := CreateTemplateCache(); err != nil {
			return err
		}
	}

	templateObj, ok := app.TemplateCache[tmpl]
	if !ok {
		return errors.New("could not find the given template from template cache")
	}

	buf := new(bytes.Buffer)

	var templateData *models.TemplateData

	if err := templateObj.Execute(buf, templateData); err != nil {
		return err
	}

	if _, err := buf.WriteTo(ctx.Writer); err != nil {
		return err
	}

	return nil
}

func CreateTemplateCache() (err error) {

	app.TemplateCache = make(map[string]*template.Template)

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		templateSet, err := template.New(name).Funcs(templateFuncMap).ParseFiles(page)
		if err != nil {
			return err
		}

		// look for layouts
		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return err
		}

		// if any layouts were found
		if len(layouts) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return err
			}
		}

		app.TemplateCache[name] = templateSet
	}

	return nil
}
