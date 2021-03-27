package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/arthurkasper/DinosaurApp/api/handlers"
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

	//código que será executado a cada request, aqui podemos colocar logs, validação de cabeçalhos e etc
	middleware := negroni.New(
		negroni.NewLogger(),
	)

	handlers.MakeDonisaurHandler(router, middleware, service)

	/*router.Handle("/v1/dinosaur", middleware.With(
		negroni.Wrap(hello(service)),
	)).Methods("GET", "OPTIONS")*/

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
