package main

type miniTeam struct {
	Id           int    `json:"id"`
	Abbreviation string `json:"abbreviation"`
	Color        string `json:"color"`
	Logo         string `json:"logo"`
	Points       int    `json:"points"`
}

type GameExt struct {
	ID       int      `json:"id"`
	HomeTeam miniTeam `json:"home_team"`
	AwayTeam miniTeam `json:"away_team"`
}
