package main

import (
	"database/sql"
	"fmt"
)

func selectConferenceYear(db *sql.DB, year int, conference string) []Record {

	results, err := db.Query("SELECT * FROM teams WHERE conference=? AND recYear=?", conference, year)

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

	stmt, err := db.Prepare("INSERT INTO teams (recYear, team, conference, division, expectedWins, " +
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
