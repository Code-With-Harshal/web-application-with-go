package renderer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// region RenderTemplate
func RenderTemplate(rw http.ResponseWriter, tmpl string) {
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		println(err)
	}

	_, err = buf.WriteTo(rw)
	if err != nil {
		fmt.Println(err)

	}
}

func createTemplateCache() (map[string]*template.Template, error) {
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
