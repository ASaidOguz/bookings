package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ASaidOguz/bookings/pkg/config"
	"github.com/ASaidOguz/bookings/pkg/handlers"
	"github.com/ASaidOguz/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig

const portNumber = ":8080"

var session *scs.SessionManager

func main() {

	// change this to true when in production

	app.InProduction = false
	// if the session below has ":"
	//the fault of variable shadowing would occur !!!
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	// development mode element
	//app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)
	// Handlers.Repo because home function has receiver of repository
	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
