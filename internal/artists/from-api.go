package artists

import (
	"net/url"

	"github.com/ArtyomKr/go-mbrainz-parser/internal/client"
	"github.com/ArtyomKr/go-mbrainz-parser/models"
)

func GetArtistFromApi(c *client.Client, ID string) (models.ArtistAPIResponse, error) {
	var artistRes models.ArtistAPIResponse
	query := url.Values{"inc": {"url-rels"}}
	err := c.Fetch("GET", "/ws/2/artist/"+ID+"?inc=url-rels", "", query, nil, &artistRes)

	return artistRes, err
}
