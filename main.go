package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manuelmaciel/go-api-gorm/database"
)

func main() {

	database.DBConnection()

	r := mux.NewRouter()

	// writer and reader
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// listen and serve - requires port and router
	http.ListenAndServe(":3000", r)
}
