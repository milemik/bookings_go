package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/milemik/bookings_go/internal/config"
	"github.com/milemik/bookings_go/internal/forms"
	"github.com/milemik/bookings_go/internal/model"
	"github.com/milemik/bookings_go/internal/render"
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
	var emptyReservation model.Reservation

	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &model.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// Reservation post reservation handle of posting reservation form
func (m *Reposatory) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err == nil {
		log.Println(err)
	}
	reservation := model.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
	}
	form := forms.New(r.PostForm)

	// form.Has("first_name", r)
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation.page.tmpl", &model.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
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

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles respons of availability and sends json response
func (m *Reposatory) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	response := jsonResponse{Ok: true, Message: "Available"}
	out, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact renders search availability page
func (m *Reposatory) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &model.TemplateData{})
}
