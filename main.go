package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/m0thm4n/spotify-bot/automod"
	"github.com/m0thm4n/spotify-bot/bot"
	"github.com/m0thm4n/spotify-bot/commands"
	"github.com/m0thm4n/spotify-bot/config"
	logger "github.com/m0thm4n/spotify-bot/logger"

	"github.com/bwmarrin/discordgo"
)

var db *sql.DB
var err error
var t0 time.Time
var userMap = make(map[uint64]string)

func main() {
	// scripts.CreateNewContainer("steamcmd/steamcmd:latest", "")

	config.ReadConfig()

	discordBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := discordBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	bot.BotID = user.ID

	discordBot.AddHandler(messageCreate)
	discordBot.AddHandler(messageReactionAdd)

	err = discordBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")

	defer discordBot.Close()

	<-make(chan struct{})
	return
}

//func GuildMemberUpdate()
// func OnGuildMemberAdd(s *discordgo.Session, g *discordgo.GuildMemberAdd) {
// 	if s == nil || g == nil {
// 		return
// 	}

// 	var user = g.User
// 	if user.ID == BotID {
// 		return
// 	}

// 	st, err := s.UserChannelCreate(user.ID)
// 	if err != nil {
// 		return
// 	}
// }

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if s == nil || m == nil {
		return
	}

	if m.Author.ID == bot.BotID {
		return
	}

	if m.Content == "" {
		return
	}

	if m.Content[0] == '!' && strings.Count(m.Content, "!") < 2 {

		commands.ExecuteCommand(s, m.Message, t0)
		return
	}

	if automod.IsWordCensored(m.Message, db) {
		err := s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
		if err != nil {
			logger.WriteError("Issue deleting a censored message.", err)
		}
		return
	}

	if automod.IsWordOnTimer(m.Message, db) {
		timer := time.NewTimer(time.Minute)
		go func() {
			<-timer.C
			err := s.ChannelMessageDelete(m.ChannelID, m.ID)
			if err != nil {
				logger.WriteError("Issue deleting a timed message.", err)
				return
			}
		}()
		return
	}

}

func messageReactionAdd(s *discordgo.Session, reactMsg *discordgo.MessageReactionAdd) {
	_, err := s.ChannelMessage(reactMsg.ChannelID, reactMsg.MessageID)
	if err != nil {
		logger.WriteError("A problem occurred while getting a message.", err)
	}
}
