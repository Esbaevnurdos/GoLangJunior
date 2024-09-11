package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Esbaevnurdos/buildingapp/pkg/config"
	"github.com/Esbaevnurdos/buildingapp/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
    app = a 
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
    var ts map[string]*template.Template
    if app.UseCache {
        ts = app.TemplateCache
    } else {
        ts, _ = CreateTemplateCache()
    }

    t, ok := ts[tmpl]
    if !ok {
        log.Fatal("Could not get template from template cache")
    }

    buf := new(bytes.Buffer)

    // Make sure to pass td (template data) here
    err := t.Execute(buf, td)  
    if err != nil {
        log.Println(err)
    }

    _, err = buf.WriteTo(w)
    if err != nil {
        log.Println(err)
    }
}



func CreateTemplateCache() (map[string]*template.Template, error) {
    // myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
	    if err != nil {
		return myCache, err
	}

	   matches, err := filepath.Glob("./templates/*.layout.html")
	   	    if err != nil {
		return myCache, err
	}

	if len(matches) > 0 {
		ts, err = ts.ParseGlob("./templates/*.layout.html")
			   	    if err != nil {
		return myCache, err
	}

	}
	 myCache[name] = ts
	}

	return myCache, nil
}

