package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/arthurkasper/DinosaurApp/core/dinosaur"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	db, err := sql.Open("sqlite3", "../data/dinosaur.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	service := dinosaur.NewService(db)
	router := mux.NewRouter()
	middleware := negroni.New(
		negroni.NewLogger(),
	)

	router.Handle("/v1/dinosaur", middleware.With(
		negroni.Wrap(hello(service)),
	)).Methods("GET", "OPTIONS")

	http.Handle("/", router)

	server := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":4000",
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func hello(service dinosaur.OperationService) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		all, _ := service.GetAll()
		for _, i := range all {
			fmt.Println("service: ", i)
		}
		/*saved, err := service.Get(1)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println(saved)*/
	})
}
