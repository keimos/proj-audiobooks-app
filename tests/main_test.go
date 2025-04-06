package main

import (
	"os"
	"testing"
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
