package main

import (
	"html/template"
	"net/http"
)

type Item struct {
	Name string
	Ep   int
}

func NewEp(w http.ResponseWriter, req *http.Request) {
	items := make([]Item, 2)
	items[0] = Item{"test", 0}
	items[1] = Item{"test2", 1}
	tmpl := template.Must(template.ParseFiles("view/newEp.html"))
	tmpl.Execute(w, items)
}

func main() {
	http.HandleFunc("/", NewEp)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":5000", nil)
}
