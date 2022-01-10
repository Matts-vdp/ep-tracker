package main

import (
	"html/template"
	"net/http"
	"os"

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

/*func ChangeEp(w http.ResponseWriter, req *http.Request) {
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
}*/

func main() {
	storage.Init()
	defer storage.Close()
	port := os.Getenv("PORT")
	http.HandleFunc("/", NewEp)
	http.HandleFunc("/list", ListEp)
	//http.HandleFunc("/epchange", ChangeEp)
	http.HandleFunc("/add", AddEp)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":"+port, nil)
}
