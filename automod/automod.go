package automod

import (
	"strings"

	"database/sql"
	"fmt"

	"github.com/m0thm4n/spotify-bot/database"

	"github.com/bwmarrin/discordgo"
)

var removeableWordsMap = make(map[int]string)
var censoredWordsMap = make(map[int]string)

// ReloadTables loads the tables
func ReloadTables() {
	for k := range removeableWordsMap {
		delete(removeableWordsMap, k)
	}

	for k := range censoredWordsMap {
		delete(censoredWordsMap, k)
	}

	fmt.Println("[INFO] Word cache cleared.")
}

// LoadAutomodTables loads the tables
func LoadAutomodTables(db *sql.DB) bool {
	fmt.Println("[INFO]Loading censored words table...")
	if ok, _ := database.LoadDatabaseCensoredWords(db, &censoredWordsMap); !ok {
		return false
	}

	fmt.Println("[INFO]Loading removeable words table...")
	if ok, _ := database.LoadDatabaseTimers(db, &removeableWordsMap); !ok {
		return false
	}

	return true
}

//IsWordCensored (* discordgo.Message) bool
//Words that match this check are immediately removed from chat
func IsWordCensored(m *discordgo.Message, db *sql.DB) bool {
	tokens := strings.Split(m.Content, " ")
	for i, v := range censoredWordsMap {
		for j := 0; j < len(tokens); j++ {
			if _, ok := censoredWordsMap[i]; !ok {
				fmt.Println("[ERROR] Attempt to access index out of bounds during censor search")
				return false
			}

			if strings.EqualFold(v, tokens[j]) {
				fmt.Printf("\n[LOG] Message erased: %s", m.Content)
				return true
			}
		}
	}
	return false
}

// IsWordOnTimer loads words on timer
func IsWordOnTimer(m *discordgo.Message, db *sql.DB) bool {
	tokens := strings.Split(m.Content, " ")
	for i, v := range removeableWordsMap {
		for j := 0; j < len(tokens); j++ {
			if _, ok := removeableWordsMap[i]; !ok {
				fmt.Println("[ERROR] Attempt to access index out of bounds during removable search")
				return false
			}

			if strings.EqualFold(v, tokens[j]) {
				fmt.Printf("\n[LOG] Message queued to be erased: %s", m.Content)
				return true
			}
		}
	}
	return false
}
