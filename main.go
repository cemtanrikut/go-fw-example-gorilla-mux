package main

import (
	"container/list"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var dataList list.List

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", MainHandler).Methods(http.MethodGet)
	router.HandleFunc("/health", Health).Methods(http.MethodGet)
	router.HandleFunc("/create", CreateHandler).Methods(http.MethodPost)

	subrouter := router.PathPrefix("/products").Subrouter()
	subrouter.HandleFunc("/", MainHandler).Methods(http.MethodGet)
	// "/products/{key}/"
	subrouter.HandleFunc("/{key}/", GetHandler).Methods(http.MethodGet)

	http.Handle("/", router)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome endpoint"))
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	w.WriteHeader(http.StatusOK)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	dataList.PushBack(vars["Anything"])
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create handler routed succesfully."))
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)["key"]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(param))
}
