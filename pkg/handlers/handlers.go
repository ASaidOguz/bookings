package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ASaidOguz/bookings/pkg/config"
	"github.com/ASaidOguz/bookings/pkg/models"
	"github.com/ASaidOguz/bookings/pkg/render"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

//About is the the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again testing testing dev mod on ? "

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	// send the data to the template
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

//Availability is the the search-availability  page handler
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

//PostAvailability is the the search-availability  page handler
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Start date is :%s ,End date is :%s...", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

//AvailabilityJSON handles request for availability and sends JSON responds
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      false,
		Message: "Available!",
	}
	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

//Generals is the the Generals room page handler
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})
}

//Majors is the the Majors room page handler
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, r, "majors.page.html", &models.TemplateData{})
}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}

//Reservation is the make-reservation page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{})
}
