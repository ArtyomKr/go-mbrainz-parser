package artists

import (
	"fmt"
	"log"
	"sync"

	"github.com/ArtyomKr/go-mbrainz-parser/internal/client"
	"github.com/ArtyomKr/go-mbrainz-parser/internal/filesystem"
	"github.com/ArtyomKr/go-mbrainz-parser/models"
)

func ProcessDBArtists(artistsFromDB []models.ArtistFromDB, batchSize int, filename string) []models.ArtistAPIResponse {
	results := make([]models.ArtistAPIResponse, 0, len(artistsFromDB))
	client := client.NewClient()

	// Process in batches
	for i := 0; i < len(artistsFromDB); i += batchSize {
		end := i + batchSize
		if end > len(artistsFromDB) {
			end = len(artistsFromDB)
		}

		batch := artistsFromDB[i:end]
		batchResults := make(chan models.ArtistAPIResponse, len(batch))
		batchErrors := make(chan error, len(batch))
		var wg sync.WaitGroup

		// Process current batch
		for j, artist := range batch {
			wg.Add(1)
			go func(j int, artist models.ArtistFromDB) {
				defer wg.Done()

				fmt.Printf("Artists # %v of %v \n", i+j, len(artistsFromDB))

				res, err := GetArtistFromApi(client, artist.MBID)
				if err != nil {
					batchErrors <- fmt.Errorf("an error occurred for artist %v: %v", artist.ID, err)
					return
				}
				batchResults <- res
			}(j, artist)
		}

		wg.Wait()

		close(batchResults)
		close(batchErrors)

		// Process errors for current batch
		for err := range batchErrors {
			fmt.Println(err)
		}

		// Collect results for current batch
		for res := range batchResults {
			results = append(results, res)
		}

		err := filesystem.SaveToFile(filename, results)
		if err != nil {
			log.Fatalf("Failed to write file: %v", err)
		}
	}

	return results
}
