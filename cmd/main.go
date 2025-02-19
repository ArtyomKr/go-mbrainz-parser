package main

import (
	"fmt"
	"log"

	"github.com/ArtyomKr/go-mbrainz-parser/internal/artists"
	"github.com/ArtyomKr/go-mbrainz-parser/internal/database"
	"github.com/ArtyomKr/go-mbrainz-parser/internal/filesystem"
)

func main() {
	// Database connection string
	connStr := "user=musicbrainz dbname=musicbrainz_db sslmode=disable password=musicbrainz"

	// Initialize the database connection
	err := database.InitDB(connStr)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDB()

	// Query the database
	artist, err := artists.GetAllArtists(database.DB)
	if err != nil {
		log.Fatalf("Failed to get users: %v", err)
	}

	err = filesystem.SaveToFile("atists11.json", artist)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	fmt.Println("Data saved to artists.json")
}
