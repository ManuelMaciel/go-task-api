package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manuelmaciel/go-api-gorm/database"
	"github.com/manuelmaciel/go-api-gorm/models"
)

func main() {

	database.DBConnection()

	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Task{})

	r := mux.NewRouter()

	// writer and reader
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// listen and serve - requires port and router
	http.ListenAndServe(":3000", r)
}
