package database

import (
	"database/sql"
	"fmt"
)

// BotInfo states status and last
type BotInfo struct {
	status int
	last   string
}

const loadDatabaseTimersSQL = "SELECT * FROM TIMER_WORDS"
const loadDatabaseUsersSQL = "SELECT * FROM USERS"
const loadDatabaseCensorSQL = "SELECT * FROM CENSOR_WORDS"
const loadDatabaseBotInfoSQL = "SELECT * FROM BOT_INFO"

// Connect is used to connect the DB
func Connect() *sql.DB {
	Db, err := sql.Open("mysql", "root:root@/alfredbot")
	if err != nil {
		fmt.Println("[ERROR] Unable to connect to database: ", err)
		return nil
	}

	fmt.Println("[INFO] Connected to database.")
	return Db
}

// LoadDatabaseTimers loads the database timers
func LoadDatabaseTimers(Db *sql.DB, m *map[int]string) (bool, error) {
	fmt.Println("[INFO] Loading Removable Words...")
	rows, err := Db.Query(loadDatabaseTimersSQL)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var word string
		err = rows.Scan(&id, &word)
		if err != nil {
			return false, err
		}

		(*m)[id] = word
	}

	fmt.Println("[INFO] Removable Words loaded.")
	return true, nil
}

//LoadDatabaseUsers loads the users
func LoadDatabaseUsers(Db *sql.DB, m *map[uint64]string) (bool, error) {
	rows, err := Db.Query(loadDatabaseUsersSQL)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint64
		var word string
		err = rows.Scan(&id, &word)
		if err != nil {
			return false, err
		}

		(*m)[id] = word
	}

	fmt.Println("[INFO] Users loaded.")
	return true, nil
}

// LoadDatabaseCensoredWords cenosores the words
func LoadDatabaseCensoredWords(Db *sql.DB, m *map[int]string) (bool, error) {
	rows, err := Db.Query(loadDatabaseCensorSQL)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var word string
		err = rows.Scan(&id, &word)
		if err != nil {
			return false, err
		}

		(*m)[id] = word
	}

	fmt.Println("[INFO] Censored Words loaded.")
	return true, nil
}

// LoadBotInfo Loads the bot info
func LoadBotInfo(Db *sql.DB) (bool, BotInfo, error) {

	var info BotInfo

	rows, err := Db.Query(loadDatabaseBotInfoSQL)
	if err != nil {
		return false, info, err
	}

	defer rows.Close()

	for rows.Next() {
		var lastPlaying string
		var status int
		err = rows.Scan(&lastPlaying, &status)

		if err != nil {
			return false, info, err
		}

		info.status = status
		info.last = lastPlaying
	}

	fmt.Println("[INFO] Bot Info loaded.")
	return true, info, nil

}
