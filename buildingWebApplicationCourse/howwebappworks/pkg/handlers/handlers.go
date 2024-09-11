package handlers

import (
	"net/http"

	"github.com/Esbaevnurdos/buildingapp/pkg/config"
	"github.com/Esbaevnurdos/buildingapp/pkg/models"
	"github.com/Esbaevnurdos/buildingapp/pkg/render"
)

//
var Repo *Repository

// Repository type
type Repository struct {
	App *config.AppConfig
}

// TempleteData holds data sent from handlers to templates

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository {
		App: a,

	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
     Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
    remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

    render.RenderTemplate(w, "home.page.html", &models.TemplateData{}) // Fixed typo here
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
    stringMap := make(map[string]string)    // Ensure map is initialized
    stringMap["test"] = "Hello, again."     // Add data to the map


    remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp


    render.RenderTemplate(w, "about.page.html", &models.TemplateData{
        StringMap: stringMap,                // Pass the map to TemplateData
    })
}






// addValues adds to integers and return the sum


