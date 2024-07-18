package main

import (
	"daniel_thewindmills/templates"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {

	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var homePage templ.Component = templates.HomePage()
		var index templ.Component = templates.Index("daniel.thewindmills.com.au", homePage)

		fmt.Println("got a request!")
		templ.Handler(index).ServeHTTP(w, r)

	})

	mux.HandleFunc("/sidebar", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println(r.Form)

		var open bool = false

		if r.Form.Has("sidebar-button") {
			open = r.Form.Get("sidebar-button") == "open"
		}

		sidebar := templates.SideBar(open)
		fmt.Println("Asked for sidebar change")
		templ.Handler(sidebar).ServeHTTP(w, r)
	})

	fmt.Println("Listening on port 3000")

	http.ListenAndServe(":3000", mux)
}
