package main

import (
	"log"

	"github.com/ArtyomKr/go-mbrainz-parser/internal/artists"
	"github.com/ArtyomKr/go-mbrainz-parser/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load env variables: %v", err)
	}

	// Database connection string
	connStr := "user=musicbrainz dbname=musicbrainz_db sslmode=disable password=musicbrainz"

	// Initialize the database connection
	err = database.InitDB(connStr)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDB()

	// Query the database
	artistsFromDB, err := artists.GetAllArtists(database.DB)
	if err != nil {
		log.Fatalf("Failed to get users: %v", err)
	}

	artists.ProcessDBArtists(artistsFromDB, 1000, "artists2.json")
}
