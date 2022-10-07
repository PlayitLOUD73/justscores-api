package main

type StadiumLocation struct {
	VenueID         int     `json:"venue_id"`
	Name            string  `json:"name"`
	City            string  `json:"city"`
	State           string  `json:"state"`
	Zip             string  `json:"zip"`
	CountryCode     string  `json:"country_code"`
	Timezone        string  `json:"timezone"`
	Latitude        float32 `json:"latitude"`
	Longitude       float32 `json:"longitude"`
	Elevation       string  `json:"elevation"`
	Capacity        int     `json:"capacity"`
	YearConstructed int     `json:"year_constructed"`
	Grass           bool    `json:"grass"`
	Dome            bool    `json:"dome"`
}

type Team struct {
	ID           int             `json:"id"`
	School       string          `json:"school"`
	Mascot       string          `json:"mascot"`
	Abbreviation string          `json:"abbreviation"`
	AltName1     string          `json:"alt_name1"`
	AltName2     string          `json:"alt_name2"`
	AltName3     string          `json:"alt_name3"`
	Conference   string          `json:"conference"`
	Division     string          `json:"division"`
	Color        string          `json:"color"`
	AltColor     string          `json:"alt_color"`
	Logos        [2]string       `json:"logos"`
	Twitter      string          `json:"twitter"`
	Location     StadiumLocation `json:"location"`
}
