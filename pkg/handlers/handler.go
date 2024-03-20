package handlers

import (
	"github.com/Code-With-Harshal/web-application-with-go/pkg/renderer"
	"net/http"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	renderer.RenderTemplate(rw, "home.page.gohtml")
}

func About(rw http.ResponseWriter, r *http.Request) {
	renderer.RenderTemplate(rw, "about.page.gohtml")
}
