package rest

import (
	"github.com/gorilla/mux"
	"github.com/paweljw/cloud-native-go-kvs/pkg/kvs"
	"io"
	"log"
	"net/http"
)

func keyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, ok := vars["key"]
	if !ok {
		http.Error(w, "key must be present in request", http.StatusUnprocessableEntity)
		return
	}

	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = kvs.Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func keyValueGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, ok := vars["key"]
	if !ok {
		http.Error(w, "key must be present in request", http.StatusUnprocessableEntity)
		return
	}

	result, err := kvs.Get(key)
	if err == kvs.ErrorNoSuchKey {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(result))
	if err != nil {
		log.Println(err)
	}
}

func keyValueDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, ok := vars["key"]
	if !ok {
		http.Error(w, "key must be present in request", http.StatusUnprocessableEntity)
		return
	}

	err := kvs.Del(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func StartRestService() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/keys/{key}", keyValuePutHandler).Methods("PUT")
	r.HandleFunc("/v1/keys/{key}", keyValueGetHandler).Methods("GET")
	r.HandleFunc("/v1/keys/{key}", keyValueDeleteHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":2137", r))
}
