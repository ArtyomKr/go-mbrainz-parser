package models

import "github.com/guregu/null/v5"

type ArtistFromDB struct {
	ID       int         `json:"id"`
	MBID     string      `json:"mbid"`
	Name     string      `json:"name"`
	SortName string      `json:"sortName"`
	ISNI     null.String `json:"ISNI"`
	Comment  string      `json:"comment"`
}


