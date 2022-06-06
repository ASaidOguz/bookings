package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ASaidOguz/bookings/internal/config"
	"github.com/ASaidOguz/bookings/internal/driver"
	"github.com/ASaidOguz/bookings/internal/handlers"
	"github.com/ASaidOguz/bookings/internal/helpers"
	"github.com/ASaidOguz/bookings/internal/models"
	"github.com/ASaidOguz/bookings/internal/render"

	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig

const portNumber = ":8080"

var session *scs.SessionManager

var infoLog *log.Logger

var errorLog *log.Logger

func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	// When the main function stops db will be closed thanks to defer function we implemented
	defer db.SQL.Close()
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

func run() (*driver.DB, error) {
	//what am i going to put in session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	//infolog and error log can be print out in terminal
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// app.InProduction change this to true when in production
	app.InProduction = false
	// if the session below has ":"
	//the fault of variable shadowing would occur !!!
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session
	// Connecting to database

	log.Println("Connecting to database ....")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=2061040215")
	if err != nil {
		log.Fatal("Cannot connect to database Dying ...")
	}
	log.Println("Connected to Database...")
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc
	// development mode element
	//app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)
	return db, nil
}
