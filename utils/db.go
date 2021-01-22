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
