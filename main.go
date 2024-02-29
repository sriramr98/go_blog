package main

import (
	"log"
	"net/http"
    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

func main() {

	connStr := "postgresql://postgres:faAA2ED5ac3Be44A*D41d3db2g5C632B@roundhouse.proxy.rlwy.net:30380/railway"
    
    db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Service Alive"))
	})

	http.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		handler := BlogHandler{Db: db}
		switch r.Method {
		case "GET":
			handler.HandleGetAll(w, r)
		case "POST":
			handler.HandlePost(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	if err = http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	} else {
        log.Println("Listening on port 8080")
    }
}
