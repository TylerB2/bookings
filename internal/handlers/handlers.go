package handlers

import (
	"bookings/internal/config"
	"bookings/internal/forms"
	"bookings/internal/models"
	"bookings/internal/render"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// is the repository type
type Repository struct {
	App *config.AppConfig
}

// used by repository handlers
var Repo *Repository

// Creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

// Home Handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About page Handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "This is a string Map"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	//send it to the template
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

//Generals Suite Handler

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Renders the make Reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {

	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// Posts a reservation
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	//parse Form Data
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	//added form data to models
	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
	}

	form := forms.New(r.PostForm)
	form.Has("first_name", r)
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

}

// Renders Search Availability PAGE
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// Post Availability Search Availability PAGE
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	//Getting data from template form
	start := r.Form.Get("start")
	end := r.Form.Get("start")

	//printing
	w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))

}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

// Handles Request for Availability and sends json back
func (m *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Ok:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "   ")
	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

// Renders Search Availability PAGE
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
