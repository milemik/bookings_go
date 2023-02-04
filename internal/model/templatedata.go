package model

import "github.com/milemik/bookings_go/internal/forms"

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flesh     string
	Warning   string
	Error     string
	Form      *forms.Form
}
