package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Esbaevnurdos/buildingapp/pkg/config"
	"github.com/Esbaevnurdos/buildingapp/pkg/handlers"
	"github.com/Esbaevnurdos/buildingapp/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager


func main() {

    // change this in true when in production
    app.InProduction = false

    session = scs.New()
    session.Lifetime = 24 * time.Hour
    session.Cookie.Persist = true
    session.Cookie.SameSite = http.SameSiteLaxMode
    session.Cookie.Secure = false

    app.Session = session

    ts, err := render.CreateTemplateCache()
    if err != nil {
        log.Fatal("cannot create template cache")
    }

    app.TemplateCache = ts
    app.UseCache = false

    repo := handlers.NewRepo(&app)
    handlers.NewHandlers(repo)
    render.NewTemplates(&app)

    fmt.Printf("Starting application on port %s\n", portNumber)

    srv := &http.Server{
        Addr:    portNumber,
        Handler: routes(&app),
    }

    err = srv.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}
