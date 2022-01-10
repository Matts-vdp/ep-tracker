package storage

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Item struct {
	Id   int
	Name string
	Ep   int
}

func get(query string) []Item {
	rows, err := db.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()
	items := make([]Item, 0)
	for rows.Next() {
		var item Item
		rows.Scan(&item)
		items = append(items, item)
	}
	return items
}

func GetNew() []Item {
	query := "SELECT * FROM episodes WHERE ep = 0"
	return get(query)
}

func GetOld() []Item {
	query := "SELECT * FROM episodes WHERE ep != 0"
	return get(query)
}

func Init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	db.Exec("create table IF NOT EXISTS episodes (Id serial primary key, Name varchar(), Ep int)")
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
}

func Close() {
	db.Close()
}

func Add(name string) {
	_, err := db.Exec("insert into episodes (name, ep) values ($1, 0)", name)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
