package main

type WinList struct {
	Games  int `json:"games"`
	Wins   int `json:"wins"`
	Losses int `json:"losses"`
	Ties   int `json:"ties"`
}

type Record struct {
	Id              int     `json:"id"`
	Year            int     `json:"year"`
	Team            string  `json:"team"`
	Conference      string  `json:"conference"`
	Division        string  `json:"division"`
	ExpectedWins    float32 `json:"expectedWins"`
	Total           WinList `json:"total"`
	ConferenceGames WinList `json:"conferenceGames"`
	HomeGames       WinList `json:"homeGames"`
	AwayGames       WinList `json:"awayGames"`
}
