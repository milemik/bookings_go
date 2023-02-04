package main

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/milemik/bookings_go/internal/config"
)

func TestRouters(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing OK
	default:
		t.Errorf("type of v is not *chi.Mux, type is %T", v)
	}
}
