package renderer

import (
	"bytes"
	"fmt"
	"github.com/Code-With-Harshal/web-application-with-go/pkg/config"
	"github.com/Code-With-Harshal/web-application-with-go/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// region RenderTemplate
func RenderTemplate(rw http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Not able to get template from template cache.")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(rw)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.gohtml")

	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts

	}
	return myCache, nil
}

// endregion RenderTemplate

// region RenderTemplateBasic
//func RenderTemplateBasic(rw http.ResponseWriter, tmpl string) {
//	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.gohtml")
//	err := parsedTemplate.Execute(rw, nil)
//	if err != nil {
//		fmt.Println("error parsing template:", err)
//
//	}
//}
// endregion RenderTemplateBasic

// region RenderTemplateWithMap
//var tc = make(map[string]*template.Template)
//
//func RenderTemplateWithMap(wr http.ResponseWriter, t string) {
//	var tmpl *template.Template
//	var err error
//
//	_, inMap := tc[t]
//	if !inMap {
//		fmt.Println("Template Created and Added to catch")
//		err = createTemplateCache(t)
//		if err != nil {
//			log.Println(err)
//		}
//	} else {
//		fmt.Println("Using cached template")
//	}
//
//	tmpl = tc[t]
//	err = tmpl.Execute(wr, nil)
//
//	if err != nil {
//		fmt.Println("error parsing template:", err)
//
//	}
//}
//
//func createTemplateCache(t string) error {
//	templates := []string{
//		fmt.Sprintf("./templates/%s", t),
//		"./templates/base.layout.gohtml",
//	}
//
//	tmpl, err := template.ParseFiles(templates...)
//	if err != nil {
//		return err
//	}
//	tc[t] = tmpl
//	return nil
//}
// endregion RenderTemplateWithMap
