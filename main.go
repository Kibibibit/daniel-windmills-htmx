package main

import (
	"daniel_thewindmills/templates"
	"fmt"
	"net/http"
	"net/url"

	"github.com/a-h/templ"
)

func getCurrentPage(r *http.Request) string {
		referer, err := url.Parse(r.Header.Get("referer"))
		if (err != nil) {
				println(err)
		}
		return referer.Path
}

func switchPage(w http.ResponseWriter, r *http.Request, page templ.Component, currentPage string) {
		index := templates.Index("daniel.thewindmills.com.au", page, currentPage)
		templ.Handler(index).ServeHTTP(w,r)
}

func main() {
    

	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switchPage(w,r, templates.HomePage(), "/")
	})

	mux.HandleFunc("/showcase", func(w http.ResponseWriter, r *http.Request) {
		switchPage(w,r, templates.ShowcasePage(), "/showcase")
	})
	
	mux.HandleFunc("/sidebar/open", func(w http.ResponseWriter, r *http.Request) {
		sidebar := templates.Sidebar(true, getCurrentPage(r))
		templ.Handler(sidebar).ServeHTTP(w, r)
	})


	mux.HandleFunc("/sidebar/close", func(w http.ResponseWriter, r *http.Request) {
		sidebar := templates.Sidebar(false, getCurrentPage(r))
		templ.Handler(sidebar).ServeHTTP(w, r)
	})


	fmt.Println("Listening on port 3000")

	http.ListenAndServe(":3000", mux)
}
