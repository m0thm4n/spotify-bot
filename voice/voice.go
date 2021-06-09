package voice

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func JoinVoiceChannel(s *discordgo.Session, m *discordgo.Message) {
	voiceConnection, err := s.ChannelVoiceJoin("752394500013424640", "752394500013424644", false, true)
	if err != nil {
		logrus.Println(err)
	}

	voiceConnection.Speaking(true)
}