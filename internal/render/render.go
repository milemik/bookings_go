package render

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/justinas/nosurf"
	"github.com/milemik/bookings_go/internal/config"
	"github.com/milemik/bookings_go/internal/model"
)

var app *config.AppConfig

var pathToTemplates = "./templates"

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddTemplateData(td *model.TemplateData, r *http.Request) *model.TemplateData {
	// We can add some default data here
	td.Flesh = app.Session.PopString(r.Context(), "flesh")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")

	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *model.TemplateData) error {

	var ts map[string]*template.Template

	// Ether use cache or recreate template
	if app.UseCache {
		ts = app.TemplateCache
	} else {
		ts, _ = CreateTemplateCache()
	}

	// create template cache
	// ts = app.TemplateCache

	// get named request from cache
	t, ok := ts[tmpl]
	if !ok {
		return errors.New("could get template from cache")
	}

	// Buffer if there is some error we will use buffer to get more info on it
	buf := new(bytes.Buffer)

	td = AddTemplateData(td, r)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
		return err
	}
	// render template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
		return err
	}

	// PRIMER PRVI
	// parsedTamplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	// err := parsedTamplate.Execute(w, nil)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// Create template here
	myCache := map[string]*template.Template{}

	// get all files named *.page.tmpl from ./templates
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}
	// range true all files ending with *.page.tmpl
	for _, page := range pages {
		// filepath.Base - return last part of path
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil

}

/*
// RENDERING TEMPLATE OPTION 1
// Creating map - this map will be live as long as we have program running
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// Check if template in map
	_, inMap := tc[t]

	if !inMap {
		// We should add it to map
		log.Println("creating template and adding to cache")
		err = createTemplateCashe(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// We should render it
		log.Println("using cache template")
	}

	// get template
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCashe(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// Parse template
	parsedTemplate, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}
	tc[t] = parsedTemplate
	return nil
}
*/
