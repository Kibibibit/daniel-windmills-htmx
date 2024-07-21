package main

import (
	"daniel_thewindmills/templates"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func switchPage(w http.ResponseWriter, r *http.Request, page templ.Component) {
		index := templates.Index("daniel.thewindmills.com.au", page)
		templ.Handler(index).ServeHTTP(w,r)
}

func main() {

	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switchPage(w,r, templates.HomePage())
	})

	mux.HandleFunc("/showcase", func(w http.ResponseWriter, r *http.Request) {
		switchPage(w,r, templates.ShowcasePage())
	})
	
	mux.HandleFunc("/sidebar/open", func(w http.ResponseWriter, r *http.Request) {
		sidebar := templates.Sidebar(true)
		templ.Handler(sidebar).ServeHTTP(w, r)
	})


	mux.HandleFunc("/sidebar/close", func(w http.ResponseWriter, r *http.Request) {
		sidebar := templates.Sidebar(false)
		templ.Handler(sidebar).ServeHTTP(w, r)
	})


	fmt.Println("Listening on port 3000")

	http.ListenAndServe(":3000", mux)
}
