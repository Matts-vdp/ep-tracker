package main

import (
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/Matts-vdp/ep-tracker/data"
)

func NewEp(w http.ResponseWriter, req *http.Request) {
	items := data.GetNew()
	tmpl := template.Must(template.ParseFiles("view/newEp.html"))
	tmpl.Execute(w, items)
}

func ListEp(w http.ResponseWriter, req *http.Request) {
	items := data.GetOld()
	tmpl := template.Must(template.ParseFiles("view/listEp.html"))
	tmpl.Execute(w, items)
}

func ChangeEp(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if s := req.Form.Get("next"); s != "" {
		id, err := strconv.Atoi(s)
		if err != nil {
			return
		}
		data.UpdateItem(id, 1)
	} else if s := req.Form.Get("prev"); s != "" {
		id, err := strconv.Atoi(s)
		if err != nil {
			return
		}
		data.UpdateItem(id, -1)
	}
	ListEp(w, req)
}

func main() {
	data.Init()
	defer data.Close()
	port := os.Getenv("PORT")
	http.HandleFunc("/", NewEp)
	http.HandleFunc("/list", ListEp)
	http.HandleFunc("/epchange", ChangeEp)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":"+port, nil)
}
