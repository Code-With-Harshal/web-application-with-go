package main

import (
	"fmt"
	"github.com/Code-With-Harshal/web-application-with-go/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting listening on http://localhost%s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
