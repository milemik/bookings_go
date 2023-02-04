package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/milemik/bookings_go/internal/config"
	"github.com/milemik/bookings_go/internal/handlers"
	"github.com/milemik/bookings_go/internal/model"
	"github.com/milemik/bookings_go/internal/render"
)

// Const can't be changed in application
const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

// main is our main function
func main() {
	// Put something in session
	gob.Register(model.Reservation{})
	// change this to true if PROD
	app.InProduction = false

	session = scs.New()

	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // SSL - secure connection for PROD we should set it to TRUE

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	// For prod this should be set to true
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/foo", handlers.Foo) //this need to be fixed - cacheTemplate

	fmt.Printf("Server started on port: %s\n", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
