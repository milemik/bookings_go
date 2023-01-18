package handlers

import (
	"net/http"

	"github.com/milemik/bookings_go/pkg/config"
	"github.com/milemik/bookings_go/pkg/model"
	"github.com/milemik/bookings_go/pkg/render"
)

// Repo is reposatory used for handlers
var Repo *Reposatory

// Reposatory is reposatory type
type Reposatory struct {
	App *config.AppConfig
}

// NewRepo creates new reposatory
func NewRepo(a *config.AppConfig) *Reposatory {
	return &Reposatory{
		App: a,
	}
}

// NewHandlers sets reposatory for the handlers
func NewHandlers(r *Reposatory) {
	Repo = r
}

// Home is our home page
func (m *Reposatory) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &model.TemplateData{})
}

// About is our about page
func (m *Reposatory) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	stringMap["test"] = "Hello again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &model.TemplateData{
		StringMap: stringMap,
	})
}

// func Foo(w http.ResponseWriter, r *http.Request) {
// 	render.RenderTemplate(w, "foo.html")
// }
