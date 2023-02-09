package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/milemik/bookings_go/internal/config"
	"github.com/milemik/bookings_go/internal/model"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	// Put something in session
	gob.Register(model.Reservation{})
	// change this to true if PROD
	testApp.InProduction = false

	session = scs.New()

	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false // SSL - secure connection for PROD we should set it to TRUE

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(x int) {}
func (tw *myWriter) Write(b []byte) (int, error) {
	lenght := len(b)
	return lenght, nil
}
