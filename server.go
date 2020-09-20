package main

import (
	"database/sql"
	"fmt"
	"time"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()

	staticFileDirectory := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))

	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/food", getFoodHandler).Methods("GET")
	r.HandleFunc("/food", createFoodHandler).Methods("POST")
	r.HandleFunc("/food", deleteFoodHandler).Methods("DELETE")

	return r
}

func main() {
	// For Local Run:
	connString := "dbname=food_encyclopedia sslmode=disable"
	// For Docker Run:
	//connString := "host=food_db port=5432 user=postgres password=secret dbname=food_encyclopedia sslmode=disable"
	db, err := sql.Open("postgres", connString)

	for err != nil {
		fmt.Println("db connection failed. Retrying after 1 sec ...")
		time.Sleep(1 * time.Second)
		db, err = sql.Open("postgres", connString)
		// after 100 secods, quit the loop
	}
	
	fmt.Println("db connection OK...")
	// if err != nil {
	// 	panic(err)
	// }
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})

	r := newRouter()

	fmt.Println("Listening")
	http.ListenAndServe(":8080", r)
}
