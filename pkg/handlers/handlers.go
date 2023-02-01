package handlers

import (
	"fmt"
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
	render.RenderTemplate(w, r, "home.page.tmpl", &model.TemplateData{})
}

// About is our about page
func (m *Reposatory) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	stringMap["test"] = "Hello again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &model.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the make reservation page
func (m *Reposatory) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &model.TemplateData{})
}

// Generals renders generals room page
func (m *Reposatory) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &model.TemplateData{})
}

// Majors renders majors room page
func (m *Reposatory) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &model.TemplateData{})
}

// Availability renders search availability page
func (m *Reposatory) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &model.TemplateData{})
}

// PostAvailability renders search availability page
func (m *Reposatory) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("POSTED on search availability %s and %s", start, end)))
}

// Contact renders search availability page
func (m *Reposatory) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &model.TemplateData{})
}
