package handlers

import (
	"net/http"

	"github.com/Esbaevnurdos/buildingapp/pkg/render"
)


func Home(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w, "home.html")
}


// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w, "about.html")
}

// addValues adds to integers and return the sum


