package handlers

import (
	"net/http"

	"github.com/weigoldj/bookings/pkg/config"
	"github.com/weigoldj/bookings/pkg/models"
	"github.com/weigoldj/bookings/pkg/render"
)

// The repository used by the handlers
var Repo *Repository

// The repository which also contains the cofig
type Repository struct {
	App *config.AppConfig
}

// careate a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// set hte repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// home is a method on repository which contains our Config
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "index.page.tmpl", &models.TemplateData{})
}

// about is a method on repository which contains our Config
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "hello to the about page!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
