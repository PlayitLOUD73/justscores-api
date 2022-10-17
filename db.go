package main

import (
	"database/sql"
	"fmt"
)

func selectConferenceYear(db *sql.DB, year int, conference string) []Record {

	results, err := db.Query("SELECT * FROM records WHERE conference=? AND recYear=?", conference, year)

	var ret []Record

	if results == nil {
		return nil
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {
		var team Record
		err := results.Scan(&team.Id, &team.Year, &team.Team, &team.Conference, &team.Division, &team.ExpectedWins, &team.Total.Games,
			&team.Total.Wins, &team.Total.Losses, &team.Total.Ties, &team.ConferenceGames.Games, &team.ConferenceGames.Wins,
			&team.ConferenceGames.Losses, &team.ConferenceGames.Ties, &team.HomeGames.Games, &team.HomeGames.Wins,
			&team.HomeGames.Losses, &team.HomeGames.Ties, &team.AwayGames.Games, &team.AwayGames.Wins, &team.AwayGames.Losses,
			&team.AwayGames.Ties)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		ret = append(ret, team)
	}

	return ret

}

func insertRecord(db *sql.DB, team Record) {

	stmt, err := db.Prepare("INSERT INTO records (recYear, team, conference, division, expectedWins, " +
		"totalGames, totalWins, totalLosses, totalTies, conGames, conWins, conLosses, conTies, " +
		"homeGames, homeWins, homeLosses, homeTies, awayGames, awayWins, awayLosses, awayTies) " +
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Preparation successful for record insert.")
	}

	_, err = stmt.Exec(team.Year, team.Team, team.Conference, team.Division, team.ExpectedWins, team.Total.Games,
		team.Total.Wins, team.Total.Losses, team.Total.Ties, team.ConferenceGames.Games, team.ConferenceGames.Wins,
		team.ConferenceGames.Losses, team.ConferenceGames.Ties, team.HomeGames.Games, team.HomeGames.Wins,
		team.HomeGames.Losses, team.HomeGames.Ties, team.AwayGames.Games, team.AwayGames.Wins, team.AwayGames.Losses,
		team.AwayGames.Ties)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Entry into team db is successful")
	}
}

func selectGamesTeam(db *sql.DB, year int, team string) []Game {

	results, err := db.Query("SELECT * FROM games WHERE (hTeam=? OR aTeam=?) AND season=?", team, team, year)

	var ret []Game

	if results == nil {
		return nil
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {
		var game Game

		// hacky time
		err := results.Scan(&game.Id, &game.Season, &game.Week, &game.SeasonType, &game.StartDate, &game.StartTimeTBD,
			&game.NeutralSite, &game.ConferenceGame, &game.Attendance, &game.VenueId, &game.HomeId, &game.HomeTeam,
			&game.HomeConference, &game.HomeDivision, &game.HomePoints, &game.HomeLineScores[0],
			&game.HomeLineScores[1], &game.HomeLineScores[2], &game.HomeLineScores[3], &game.HomePostWinProb,
			&game.HomePreGameElo, &game.HomePostGameElo, &game.AwayId, &game.AwayTeam, &game.AwayConference,
			&game.AwayDivision, &game.AwayPoints, &game.AwayLineScores[0], &game.AwayLineScores[1],
			&game.AwayLineScores[2], &game.AwayLineScores[3], &game.AwayPostWinProb, &game.AwayPreGameElo,
			&game.AwayPostGameElo, &game.ExcitementIndex, &game.Highlights, &game.Notes)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		ret = append(ret, game)
	}

	return ret

}

func insertGameTeam(db *sql.DB, game Game) {

	stmt, err := db.Prepare("INSERT INTO games VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, " +
		"?, ?, ?, ?, ?, ?, ?, ?, ?, ?, " +
		"?, ?, ?, ?, ?, ?, ?, ?, ?, ?, " +
		"?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Preparation successful for game insert.")
	}

	_, err = stmt.Exec(game.Id, game.Season, game.Week, game.SeasonType, game.StartDate, game.StartTimeTBD,
		game.NeutralSite, game.ConferenceGame, game.Attendance, game.VenueId, game.HomeId, game.HomeTeam,
		game.HomeConference, game.HomeDivision, game.HomePoints, game.HomeLineScores[0],
		game.HomeLineScores[1], game.HomeLineScores[2], game.HomeLineScores[3], game.HomePostWinProb,
		game.HomePreGameElo, game.HomePostGameElo, game.AwayId, game.AwayTeam, game.AwayConference,
		game.AwayDivision, game.AwayPoints, game.AwayLineScores[0], game.AwayLineScores[1],
		game.AwayLineScores[2], game.AwayLineScores[3], game.AwayPostWinProb, game.AwayPreGameElo,
		game.AwayPostGameElo, game.ExcitementIndex, game.Highlights, game.Notes)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Entry into game db is successful")
	}
}

// very basic needs to be redone to support returning the location
func selectLocation(db *sql.DB, venueID int) bool {

	results, err := db.Query("SELECT * FROM locations WHERE venueId=?", venueID)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if results == nil {
		return false
	} else {
		return true
	}
}

func insertLocation(db *sql.DB, loc StadiumLocation) {

	if !selectLocation(db, loc.VenueID) {
		return
	}

	stmt, err := db.Prepare("INSERT INTO locations VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("Preparating successful for location insert.")
	}

	_, err = stmt.Exec(loc.VenueID, loc.Name, loc.City, loc.State, loc.Zip,
		loc.CountryCode, loc.Timezone, loc.Latitude, loc.Longitude, loc.Elevation,
		loc.Capacity, loc.YearConstructed, loc.Grass, loc.Dome)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Entry into location db is successful")
	}

}

func insertTeam(db *sql.DB, team Team) {

	insertLocation(db, team.Location)

	stmt, err := db.Prepare("INSERT INTO teams VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("Preparing successful for team insert.")
	}

	_, err = stmt.Exec(team.ID, team.School, team.Mascot, team.Abbreviation, team.AltName1,
		team.AltName2, team.AltName3, team.Conference, team.Division,
		team.Color, team.AltColor, team.Logos[0], team.Logos[1], team.Twitter,
		team.Location.VenueID)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Entry into team db is successful")
	}

}
