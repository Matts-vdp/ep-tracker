package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Matts-vdp/ep-tracker/storage"
)

func NewEp(w http.ResponseWriter, req *http.Request) {
	items := storage.GetNew()
	tmpl := template.Must(template.ParseFiles("view/newEp.html"))
	tmpl.Execute(w, items)
}

func ListEp(w http.ResponseWriter, req *http.Request) {
	items := storage.GetOld()
	tmpl := template.Must(template.ParseFiles("view/listEp.html"))
	tmpl.Execute(w, items)
}

func AddEp(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	name := req.Form.Get("name")
	storage.Add(name)
	http.Redirect(w, req, "/", http.StatusFound)
}
func DelEp(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	id, err := strconv.Atoi(req.Form.Get("id"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, req, "/", http.StatusFound)
		return
	}
	storage.Del(id)
	http.Redirect(w, req, "/", http.StatusFound)
}

func ChangeEp(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if s := req.Form.Get("next"); s != "" {
		id, err := strconv.Atoi(s)
		if err != nil {
			http.Redirect(w, req, "/list", http.StatusFound)
			return
		}
		storage.UpdateEp(id, 1)
	} else if s := req.Form.Get("prev"); s != "" {
		id, err := strconv.Atoi(s)
		if err != nil {
			return
		}
		storage.UpdateEp(id, -1)
	} else if s := req.Form.Get("nexts"); s != "" {
		id, err := strconv.Atoi(s)
		if err != nil {
			return
		}
		storage.UpdateSeason(id, 1)
	} else if s := req.Form.Get("prevs"); s != "" {
		id, err := strconv.Atoi(s)
		if err != nil {
			return
		}
		storage.UpdateSeason(id, -1)
	} else if s := req.Form.Get("del"); s != "" {
		id, err := strconv.Atoi(s)
		if err != nil {
			return
		}
		storage.Del(id)
	} else if s := req.Form.Get("done"); s != "" {
		id, err := strconv.Atoi(s)
		if err != nil {
			return
		}
		storage.Done(id)
	}
	target := req.Form.Get("target")
	http.Redirect(w, req, target, http.StatusFound)
}

func main() {
	storage.Init()
	defer storage.Close()
	port := os.Getenv("PORT")
	http.HandleFunc("/", NewEp)
	http.HandleFunc("/list", ListEp)
	http.HandleFunc("/add", AddEp)
	http.HandleFunc("/epchange", ChangeEp)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":"+port, nil)
}
