package bot

import (
	"github.com/bwmarrin/discordgo"
)

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

func GetSession(token string) (*discordgo.Session, error) {
	ds, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	ds.AddHandler(onMessage)
	ds.Identify.Intents = discordgo.IntentGuildMessages
	return ds, nil
}
