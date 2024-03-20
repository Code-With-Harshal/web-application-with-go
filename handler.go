package main

import (
	"net/http"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "home.page.gohtml")
}

func About(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "about.page.gohtml")
}
