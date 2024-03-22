package handlers

import (
	"github.com/Code-With-Harshal/web-application-with-go/pkg/config"
	"github.com/Code-With-Harshal/web-application-with-go/pkg/models"
	"github.com/Code-With-Harshal/web-application-with-go/pkg/renderer"
	"net/http"
)

var Repo *Repository

// Repository Datatype
type Repository struct {
	app *config.AppConfig
}

// NewRepository creates new repository
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

// NewHandler sets repository to handlers
func NewHandler(r *Repository) {
	Repo = r
}
func (m *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.app.Session.Put(r.Context(), "remote_ip", remoteIp)
	renderer.RenderTemplate(rw, "home.page.gohtml", &models.TemplateData{})
}

func (m *Repository) About(rw http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	remoteIp := m.app.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	renderer.RenderTemplate(rw, "about.page.gohtml", &models.TemplateData{StringMap: stringMap})
}
