package main

import (
	"database/sql"
	"encurtador_url/src/internal/handler"
	"encurtador_url/src/internal/repository"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	conn, err := OpenConn()
	if err != nil {
		return
	}
	defer conn.Close()

	repo := repository.NewCodeRepository(conn)
	handler := handler.NewHandler(repo)

	http.HandleFunc("/save", handler.CreateCode)
	http.HandleFunc("/get/", handler.GetByCode)

	fmt.Println("Server is running on 5000")
	http.ListenAndServe(":5000", nil)
}

func OpenConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=codeDB sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
