package models

type LifeSpan struct {
	Begin string `json:"begin"`
	End   string `json:"end"`
	Ended bool   `json:"ended"`
}

type EndArea struct {
	Type           *string `json:"type"`
	Name           string  `json:"name"`
	ID             string  `json:"id"`
	SortName       string  `json:"sort-name"`
	TypeID         *string `json:"type-id"`
	Disambiguation string  `json:"disambiguation"`
}

type Relation struct {
	End             *string                `json:"end"`
	AttributeValues map[string]interface{} `json:"attribute-values"`
	TypeID          string                 `json:"type-id"`
	Direction       string                 `json:"direction"`
	SourceCredit    string                 `json:"source-credit"`
	AttributeIDs    map[string]string      `json:"attribute-ids"`
	Begin           *string                `json:"begin"`
	TargetType      string                 `json:"target-type"`
	Ended           bool                   `json:"ended"`
	Type            string                 `json:"type"`
	TargetCredit    string                 `json:"target-credit"`
	Attributes      []interface{}          `json:"attributes"`
	URL             struct {
		ID       string `json:"id"`
		Resource string `json:"resource"`
	} `json:"url"`
}

type Area struct {
	Name           string   `json:"name"`
	ID             string   `json:"id"`
	Type           *string  `json:"type"`
	Disambiguation string   `json:"disambiguation"`
	ISO31661Codes  []string `json:"iso-3166-1-codes"`
	TypeID         *string  `json:"type-id"`
	SortName       string   `json:"sort-name"`
}

type BeginArea struct {
	Disambiguation string  `json:"disambiguation"`
	SortName       string  `json:"sort-name"`
	TypeID         *string `json:"type-id"`
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Type           *string `json:"type"`
}

type ArtistAPIResponse struct {
	LifeSpan       LifeSpan   `json:"life-span"`
	GenderID       string     `json:"gender-id"`
	IPIs           []string   `json:"ipis"`
	Type           string     `json:"type"`
	EndArea        EndArea    `json:"end-area"`
	Isnis          []string   `json:"isnis"`
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	Relations      []Relation `json:"relations"`
	SortName       string     `json:"sort-name"`
	TypeID         string     `json:"type-id"`
	Area           Area       `json:"area"`
	Country        string     `json:"country"`
	Gender         string     `json:"gender"`
	BeginArea      BeginArea  `json:"begin-area"`
	Disambiguation string     `json:"disambiguation"`
}
