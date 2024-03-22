package main

import (
	"fmt"
	"github.com/Code-With-Harshal/web-application-with-go/pkg/config"
	"github.com/Code-With-Harshal/web-application-with-go/pkg/handlers"
	"github.com/Code-With-Harshal/web-application-with-go/pkg/renderer"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := renderer.CreateTemplateCache()
	if err != nil {
		fmt.Println("Error while creating template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepository(&app)
	handlers.NewHandler(repo)
	renderer.NewTemplate(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting listening on http://localhost%s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
