package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/FilipeParreiras/Bookings/pkg/config"
	"github.com/FilipeParreiras/Bookings/pkg/models"
	"github.com/FilipeParreiras/Bookings/pkg/render"
)

// Repo the repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Handler function
// An handler function needs to have these two arguments
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// gets remote IP Address and stores it in the session
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	// Render the template
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
// gives acess to anything inside repository
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Creates stringMap that contains key-values to use in HTML code
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	// Gets the remote IP from session and stores it in stringmap
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// Render the template
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	// Render the template
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	// Render the template
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	// Render the template
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	// Render the template
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	// Get data from user input
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

// jsonResponse is only used in AvailabilityJSON
type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles requests for availability and send JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	// Parameters to send
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	// Creates the JSON
	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	// What kind of response am i sending
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	// Render the template
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
