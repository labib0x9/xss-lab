package main

import (
	"html/template"
	"net/http"
)

// var templ = template.Must(template.ParseFiles("templates/hello.html"))

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	q := r.URL.Query().Get("name")

// 	data := template.HTML(q)

// 	templ.Execute(w, map[string]any{
// 		"Name": data,
// 	})
// }

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("http://ip:port/echo?name=Labib"))
// }

func main() {

	// http.HandleFunc("/echo", helloHandler)
	// http.HandleFunc("/", homeHandler)

	// http.ListenAndServe(":8080", nil)

	mux := http.NewServeMux()
	mux.HandleFunc("")

}
