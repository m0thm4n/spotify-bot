package general

import (
	"fmt"
	"strings"
	"time"

	"github.com/m0thm4n/spotify-bot/automod"

	"github.com/bwmarrin/discordgo"
)

//HandleInfoCommand is the !info command
func HandleInfoCommand(s *discordgo.Session, m *discordgo.Message, t0 time.Time) {

	t1 := time.Now()
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("[ERROR] Issue finding channel by ID: ", err)
		return
	}

	channelName := channel.Name
	message := "```txt\n%s\n%s\n%-16s%-20s\n%-16s%-20s\n%-16s%-20s```"
	message = fmt.Sprintf(message, "GameBot Information", strings.Repeat("-", len("GameBot Information")), "ChannelID", m.ChannelID, "Channel Name", channelName, "Uptime", (t1.Sub(t0).String()))
	s.ChannelMessageSend(m.ChannelID, message)
}

func HandleHelpCommand(s *discordgo.Session, m *discordgo.Message) {
	message := "```txt\n%s\n%s\n%-16s\t%-20s\n%-16s\t%-20s\n%-16s\t%-20s\n%-16s\t%-20s\n%-16s\t%-20s```"
	message = fmt.Sprintf(message, "Help Information", strings.Repeat("-", len("Help Information")),
		"Create a server", "!createserver <image you want the server to be> <name of server> <port for server>",
		"Delete a server", "!deleteserver <id of server here>",
		"List all servers", "!listservers",
		"Start a server", "!startserver <id of server here>",
		"Stop a server", "!stopserver <id of server here>")
	s.ChannelMessageSend(m.ChannelID, message)
}

//HandlePingCommand is for !ping
func HandlePingCommand(s *discordgo.Session, m *discordgo.Message) {

	s.ChannelMessageSend(m.ChannelID, "pong")
}

func HandleReloadCommand(s *discordgo.Session, m *discordgo.Message) {
	automod.ReloadTables()
	s.ChannelMessageSend(m.ChannelID, "Database tables have been reloaded!")
}

func HandlePurgeCommand() {

}
