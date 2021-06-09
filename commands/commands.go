package commands

import (
	"github.com/m0thm4n/spotify-bot/spotify"
	"github.com/m0thm4n/spotify-bot/voice"
	"strings"

	"github.com/m0thm4n/spotify-bot/general"

	"time"

	"github.com/bwmarrin/discordgo"
)

// Games stores the games for the game picker

//ExecuteCommand Parses and executes the command from the calling code
func ExecuteCommand(s *discordgo.Session, m *discordgo.Message, T0 time.Time) {
	msg := strings.Split(strings.TrimSpace(m.Content), "!")[1]

	if len(msg) > 2 {
		msg = strings.Split(strings.Split(m.Content, " ")[0], "!")[1]
	}

	switch msg {
	case "info":
		general.HandleInfoCommand(s, m, T0)
	case "ping":
		general.HandlePingCommand(s, m)
	case "help":
		general.HandleHelpCommand(s, m)
	case "reload":
		general.HandleReloadCommand(s, m)
	case "purge":
	case "join":
		voice.JoinVoiceChannel(s, m)
	case "play":
		spotify.Play(s, m)
	}
}

// HandleUnknownCommand is the default for any commands not listed
func HandleUnknownCommand(s *discordgo.Session, m *discordgo.Message, msg string) {

	c, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		println("Unable to open User Channel: ", err)
		return
	}
	s.ChannelMessageSend(c.ID, "The command ` "+msg+" ` is not recognized.")
}

// HandleWrongPermissions handles wrong permissions
func HandleWrongPermissions(s *discordgo.Session, m *discordgo.Message, msg string) {

	c, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		println("Unable to open User Channel: ", err)
		return
	}
	s.ChannelMessageSend(c.ID, "The command ` "+msg+" ` is not available to you.")
}
