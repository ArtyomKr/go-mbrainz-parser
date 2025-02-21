package artists

import (
	"database/sql"
	"fmt"

	"github.com/ArtyomKr/go-mbrainz-parser/models"
)

func GetAllArtists(db *sql.DB) ([]models.ArtistFromDB, error) {
	query := `
		SELECT a.id, a.gid, a.name, a.sort_name, a.comment, ai.isni
		FROM artist a
		LEFT JOIN artist_isni ai ON a.id = ai.artist
		ORDER BY a.id
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query artists: %w", err)
	}
	defer rows.Close()

	var artists []models.ArtistFromDB

	for rows.Next() {
		var artist models.ArtistFromDB
		if err := rows.Scan(&artist.ID, &artist.MBID, &artist.Name, &artist.SortName, &artist.Comment, &artist.ISNI); err != nil {
			return nil, fmt.Errorf("failed to scan artist: %w", err)
		}
		artists = append(artists, artist)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return artists, nil
}
