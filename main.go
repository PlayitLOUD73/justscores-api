package main

import (
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
)

//go:embed "secrets.txt"
var rawKey string

var PAT string

var baseURL string = "https://api.collegefootballdata.com/"
var db *sql.DB

func conferenceLongToShort(con string) string {

	if con == "Pac-12" {
		return "PAC"
	}
	if con == "Big Ten" {
		return "B1G"
	}
	if con == "Big 12" {
		return "B12"
	} else {
		return con
	}
}

func conferenceShortToLong(con string) string {
	if con == "PAC" {
		return "Pac-12"
	}
	if con == "B1G" {
		return "Big Ten"
	}
	if con == "B12" {
		return "Big 12"
	} else {
		return con
	}

}

func records(w http.ResponseWriter, r *http.Request) {
	var teamRecords []Record
	var client http.Client
	var bearer string = "Bearer " + PAT

	var year, conference string

	year = r.URL.Query().Get("year")
	//team = r.URL.Query().Get("team")
	conference = r.URL.Query().Get("conference")

	y, _ := strconv.Atoi(year)

	results := selectConferenceYear(db, y, conferenceShortToLong(conference))

	if results != nil {
		teamRecords = results
	} else {
		url := baseURL + "records" + "?conference=" + conference + "&year=" + year
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Print(err)
		}

		req.Header.Add("Authorization", bearer)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Print(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &teamRecords)
		if err != nil {
			fmt.Print(err)
			return
		}

		for i := range teamRecords {
			insertRecord(db, teamRecords[i])
		}

		results := selectConferenceYear(db, y, conferenceShortToLong(conference))
		if results != nil {
			teamRecords = results
		}
	}

	toSend, _ := json.Marshal(teamRecords)
	_, err := io.WriteString(w, string(toSend))
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print("success!")
}

func main() {

	keyInfo := strings.Split(rawKey, "=")
	PAT = keyInfo[1]

	dbs, err := sql.Open("mysql", "root:V4gAb0ND2k1@tcp(127.0.0.1:3306)/football")
	db = dbs
	if err != nil {
		panic(err.Error())
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/records", records)

	err = http.ListenAndServe(":5050", mux)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = net.Listen("tcp4", "0.0.0.0:5050")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
}
