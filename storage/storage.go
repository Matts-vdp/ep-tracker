package storage

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Item struct {
	Id     int
	Name   string
	Season int
	Ep     int
	Done   bool
}

func get(query string) []Item {
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
		return nil
	}
	defer rows.Close()
	items := make([]Item, 0)
	for rows.Next() {
		var item Item
		rows.Scan(&item.Id, &item.Name, &item.Season, &item.Ep, &item.Done)
		items = append(items, item)
	}
	return items
}

func GetNew() []Item {
	query := "SELECT * FROM episodes WHERE season=1 and ep=0 ORDER BY done ASC, id DESC"
	return get(query)
}

func GetOld() []Item {
	query := "SELECT * FROM episodes where not (ep = 0 and season = 1) ORDER BY done ASC, id DESC"
	return get(query)
}

func Init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	_, err = db.Exec("create table IF NOT EXISTS episodes (Id serial primary key, Name varchar(20), Season int,  Ep int, Done boolean)")
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
}

func Close() {
	db.Close()
}

func Add(name string) {
	_, err := db.Exec("insert into episodes (name, season, ep, done) values ($1, 1, 0, FALSE)", name)
	if err != nil {
		log.Println(err)
		return
	}
}

func Del(id int) {
	_, err := db.Exec("delete from episodes where id=$1", id)
	if err != nil {
		log.Println(err)
		return
	}
}

func Done(id int) {
	row := db.QueryRow("select done from episodes where id=$1", id)
	var done bool
	err := row.Scan(&done)
	if err != nil {
		log.Println(err)
		return
	}
	done = !done
	_, err = db.Exec("update episodes set done=$1 where id=$2", done, id)
	if err != nil {
		log.Println(err)
		return
	}
}

func UpdateEp(id, val int) {
	row := db.QueryRow("select ep from episodes where id=$1", id)
	var ep int
	err := row.Scan(&ep)
	if err != nil {
		log.Println(err)
		return
	}
	ep += val
	if ep < 1 {
		ep = 1
	}
	_, err = db.Exec("update episodes set ep=$1 where id=$2", ep, id)
	if err != nil {
		log.Println(err)
		return
	}
}
