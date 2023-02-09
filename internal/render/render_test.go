package render

import (
	"net/http"
	"testing"

	"github.com/milemik/bookings_go/internal/model"
)

func TestAddTemplateData(t *testing.T) {
	var td model.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flesh", "123")
	result := AddTemplateData(&td, r)

	if result.Flesh != "123" {
		t.Error("Flesh value is not equal to 123")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = RenderTemplate(&ww, r, "home.page.tmpl", &model.TemplateData{})

	if err != nil {
		t.Error("error writing template to browser")
	}

	err = RenderTemplate(&ww, r, "dont-exist-templ.page.tmpl", &model.TemplateData{})

	if err == nil {
		t.Error("error writing non existing template to browser")
	}
}

func TestNewTemplates(t *testing.T) {
	NewTemplates(&testApp)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
}
