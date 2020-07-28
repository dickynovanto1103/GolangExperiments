package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *handler {
	return &handler{db:db}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query("SELECT * FROM Person")
	if err != nil {
		log.Printf("error in DB query, err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("rows: %v", rows)
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Printf("error in row scan, err: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		res := fmt.Sprintf("id, name, age: %v %v %v", id, name, age)
		log.Printf(res)
		w.Write([]byte(res))
	}

}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/testDB")
	if err != nil {
		log.Fatal("error in connecting to DB:", err)
	}
	//db.SetConnMaxLifetime()
	db.SetMaxOpenConns(100)
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("error in db PING, err: %v", err)
	}

	handler := NewHandler(db)

	mux := http.NewServeMux()
	mux.Handle("/", handler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}