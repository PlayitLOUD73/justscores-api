package main

type Game struct {
	Id              int     `json:"id"`
	Season          int     `json:"season"`
	Week            int     `json:"week"`
	SeasonType      string  `json:"season_type"`
	StartDate       string  `json:"start_date"`
	StartTimeTBD    bool    `json:"start_time_tbd"`
	NeutralSite     bool    `json:"neutral_site"`
	ConferenceGame  bool    `json:"conference_game"`
	Attendance      int     `json:"attendance"`
	VenueId         int     `json:"venue_id"`
	HomeId          int     `json:"home_id"`
	HomeTeam        string  `json:"home_team"`
	HomeConference  string  `json:"home_conference"`
	HomeDivision    string  `json:"home_division"`
	HomePoints      int     `json:"home_points"`
	HomeLineScores  [4]int  `json:"home_line_scores"`
	HomePostWinProb float32 `json:"home_post_win_prob"`
	HomePreGameElo  int     `json:"home_pre_game_elo"`
	HomePostGameElo int     `json:"home_post_game_elo"`
	AwayId          int     `json:"away_id"`
	AwayTeam        string  `json:"away_team"`
	AwayConference  string  `json:"away_conference"`
	AwayDivision    string  `json:"away_division"`
	AwayPoints      int     `json:"away_points"`
	AwayLineScores  [4]int  `json:"away_line_scores"`
	AwayPostWinProb float32 `json:"away_post_win_prob"`
	AwayPreGameElo  int     `json:"away_pre_game_elo"`
	AwayPostGameElo int     `json:"away_post_game_elo"`
	ExcitementIndex float32 `json:"excitement_index"`
	Highlights      string  `json:"highlights"`
	Notes           string  `json:"notes"`
}

type SimpleGame struct {
	Id         int    `json:"id"`
	HomeTeam   string `json:"home_team"`
	HomePoints int    `json:"home_points"`
	AwayTeam   string `json:"away_team"`
	AwayPoints int    `json:"away_points"`
}
