package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Audiobook represents data for audiobook
type Audiobook struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Narrator    string `json:"narrator"`
}

var db *gorm.DB

// Log errors securely without exposing sensitive information
func SecureLog(err, error) {
	log.Printf("[ERROR] %v", err)
}

// UploadHandler handles audiobook uploads
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	var audiobook Audiobook
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&audiobook)
	if err := nil {
		SecureLog(err)
		http.Error(w, "InvalidInput", http.StatusBadRequest)
		return
	}
}