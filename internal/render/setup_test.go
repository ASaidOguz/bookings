package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/ASaidOguz/bookings/internal/config"
	"github.com/ASaidOguz/bookings/internal/models"
	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var TestApp config.AppConfig

func TestMain(m *testing.M) {
	//what am i going to put in session
	gob.Register(models.Reservation{})

	// app.InProduction change this to true when in production
	TestApp.InProduction = false
	// if the session below has ":"
	//the fault of variable shadowing would occur !!!
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	//putting false for test purposes
	session.Cookie.Secure = false

	TestApp.Session = session
	app = &TestApp
	os.Exit(m.Run())
}

type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {

}

func (tw *myWriter) Write(b []byte) (int, error) {
	length := len(b)

	return length, nil
}
