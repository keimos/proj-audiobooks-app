package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Audiobook represents data for audiobook
type Audiobook struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Narrator   string `json:"narrator"`
	Duration   int    `json:"duration"`
	UploadedAt string `json:"uploaded_at"`
}

var db *gorm.DB

// Log errors securely without exposing sensitive information
func SecureLog(err error) {
	log.Printf("[ERROR] %v", err)
}

// UploadHandler handles audiobook uploads
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	var audiobook Audiobook
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&audiobook)
	if err != nil {
		SecureLog(err)
		http.Error(w, "InvalidInput", http.StatusBadRequest)
		return
	}

	audiobook.ID = fmt.Sprintf("audiobook_%d", time.Now().UnixNano())
	audiobook.UploadedAt = time.Now().Format(time.RFC3339)

	if err := db.Create(&audiobook).Error; err != nil {
		SecureLog(err)
		http.Error(w, "DatabaseError", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(audiobook)
}

// ListAudiobooksHandler lists all uploaded audiobooks
func ListAudiobooksHandler(w http.ResponseWriter, r *http.Request) {
	var audiobooks []Audiobook
	if err := db.Find(&audiobooks).Error; err != nil {
		SecureLog(err)
		http.Error(w, "Failed to list audiobooks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(audiobooks)
}

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("audiobooks.db"), &gorm.Config{})
	if err != nil {
		SecureLog(err)
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Audiobook{})
	if err != nil {
		SecureLog(err)
		panic("failed to migrate database")
	}
}
