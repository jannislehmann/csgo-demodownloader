package utils

import (
	"database/sql"
	"log"

	// Needs a blank import.
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	localDb, err := sql.Open("sqlite3", "./configs/db.sqlite")
	if err != nil {
		log.Panic(err)
	}

	db = localDb

	sqlStmt := "CREATE TABLE IF NOT EXISTS matches (id integer not null primary key);"
	_, errExec := db.Exec(sqlStmt)
	if errExec != nil {
		log.Panic(errExec)
	}

	// We need to store share codes multiple times if they belong to different csgo users.
	sqlStmt = "CREATE TABLE IF NOT EXISTS shareCodes (steamid integer NOT NULL, shareCode varchar(34) NOT NULL, PRIMARY KEY (steamid, shareCode));"
	_, errExec = db.Exec(sqlStmt)
	if errExec != nil {
		log.Panic(errExec)
	}
}

// GetLatestShareCode returns the latest saved share code for a steam id.
// steamID should be uint64. However, this method is only called using string steamids coming from the config.
func GetLatestShareCode(steamID string) string {
	shareCode := ""
	err := db.QueryRow("SELECT shareCode FROM shareCodes WHERE steamid = ? ORDER BY rowid DESC LIMIT 1", steamID).Scan(&shareCode)
	if err != nil && err != sql.ErrNoRows {
		log.Print(err)
	}

	return shareCode
}

// AddShareCode saves a share code associated to a steam id.
func AddShareCode(steamID string, shareCode string) {
	stmt, errPrepare := db.Prepare("insert into shareCodes (steamid, shareCode) values (?, ?)")
	_, errExec := stmt.Exec(steamID, shareCode)

	if errPrepare == nil && errExec == nil {
		log.Printf("added shareCode %v to the database for steam id %v\n", shareCode, steamID)
	}
}

// CheckIfMatchExistsAlready checks whether the match id is already marked as downloaded (contained) in the database.
func CheckIfMatchExistsAlready(matchID uint64) bool {
	err := db.QueryRow("select id from matches where id = ?", matchID).Scan(&matchID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}

		return false
	}

	return true
}

// AddMatchToDatabase adds a match id to the database and marks it as already downloaded
func AddMatchToDatabase(matchID uint64) {
	if CheckIfMatchExistsAlready(matchID) {
		return
	}

	stmt, errPrepare := db.Prepare("insert into matches (id) values (?)")
	_, errExec := stmt.Exec(matchID)

	if errPrepare == nil && errExec == nil {
		log.Printf("added match %d to the database\n", matchID)
	}
}
