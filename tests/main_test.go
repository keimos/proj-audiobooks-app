package main

import (
	"os"
	"testing"

	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	os.Remove(("test.db"))
	// Initialize the database connection
	dbInit, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db = dbInit
	// Migrate the schema
	db.AutoMigrate(&Audiobook{})
	code := m.Run()
	os.Remove("test.db")
	os.Exit(code)
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/upload", UploadHandler).Methods("GET")
	router.HandleFunc("/audiobook/{id}", GetAudiobookHandler).Methods("GET")
	router.HandleFunc("/audiobooks", ListAudiobooksHandler).Methods("POST")
	return router
}
