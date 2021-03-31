package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/arthurkasper/DinosaurApp/core/dinosaur"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func MakeDonisaurHandler(router *mux.Router, n *negroni.Negroni, service dinosaur.OperationService) {
	router.Handle("/v1/dinosaur", n.With(
		negroni.Wrap(getAllDinosaur(service)),
	)).Methods("GET", "OPTIONS")

	router.Handle("/v1/dinosaur/{id}", n.With(
		negroni.Wrap(getDinosaur(service)),
	)).Methods("GET", "OPTIONS")

	router.Handle("/v1/dinosaur", n.With(
		negroni.Wrap(storeDinosaur(service)),
	)).Methods("POST", "OPTIONS")

	router.Handle("/v1/dinosaur/{id}", n.With(
		negroni.Wrap(updateDinosaur(service)),
	)).Methods("PUT", "OPTIONS")

	router.Handle("/v1/dinosaur/{id}", n.With(
		negroni.Wrap(removeDinosaur(service)),
	)).Methods("DELETE", "OPTIONS")

}

func getAllDinosaur(service dinosaur.OperationService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "aplication/json")
		allDinosaur, err := service.GetAll()
		if err != nil {
			w.Write(formatJSONerror(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(allDinosaur)
		if err != nil {
			w.Write(formatJSONerror(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	})
}

func getDinosaur(service dinosaur.OperationService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "aplication/json")

		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.Write(formatJSONerror(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		dinosaur, err := service.Get(id)
		if err != nil {
			w.Write(formatJSONerror(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(dinosaur)
		if err != nil {
			w.Write(formatJSONerror(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	})
}

func storeDinosaur(service dinosaur.OperationService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-type", "application/json")

		var dinosaur dinosaur.Dinosaur

		//get user data sending by body
		err := json.NewDecoder(r.Body).Decode(&dinosaur)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONerror(err.Error()))
			return
		}

		err = service.Store(&dinosaur)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONerror(err.Error()))
		}
		w.WriteHeader(http.StatusCreated)

	})
}

func updateDinosaur(service dinosaur.OperationService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.Write(formatJSONerror(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
		}

		var dinosaur dinosaur.Dinosaur
		dinosaur.ID = id
		err = json.NewDecoder(r.Body).Decode(&dinosaur)
		if err != nil {
			w.Write(formatJSONerror(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
		}

		err = service.Update(&dinosaur)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONerror(err.Error()))
		}

	})
}

func removeDinosaur(service dinosaur.OperationService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.Write(formatJSONerror(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
		}

		err = service.Remove(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONerror(err.Error()))
		}

	})

}
