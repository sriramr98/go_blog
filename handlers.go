package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Blog struct {
	Id         int    `db:"id"`
	Title      string `db:"title"`
	Desc       string `db:"desc"`
	Created_at string `db:"created_at"`
}

type BlogHandler struct {
	Db *sqlx.DB
}

func (h BlogHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	results := []Blog{}
	if err := h.Db.Select(&results, "SELECT * FROM blogs LIMIT 10;"); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(results); err != nil {
		log.Fatal(err)
	}
}

func (h BlogHandler) HandleGetOne(w http.ResponseWriter, r *http.Request) {
}

func (h BlogHandler) HandlePost(w http.ResponseWriter, r *http.Request) {
	var blog Blog
	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
        w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"error\": \"Invalid request body\"}"))
		return
	}

	fmt.Println(blog)

	_, err = h.Db.Exec("INSERT INTO blogs (id, title, desc, created_at) VALUES($1, $2, $3, $4);", blog.Id, blog.Title, blog.Desc, blog.Created_at)
	if err != nil {
		log.Println("Failed to insert blog")
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
        w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"error\": \"Unable to create blog due to " + err.Error() + "\"}"))
        return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(blog); err != nil {
		log.Fatal(err)
	}
}
