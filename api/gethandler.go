package api

import (
	HELPER "notification-bot/helpers"
	SERVICE "notification-bot/services"

	// 
	// REPO "notification-bot/repository"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	// Check Channel
	if HELPER.StringContain(botChannel, m.ChannelID) {
		// If Bot it self end progres
		if m.Author.ID == s.State.User.ID {
			return
		}
		// Check If someone Metion Bot => Will greating
		if len(m.Mentions) > 0 {
			for _, men := range m.Mentions {
				if men.ID == s.State.User.ID {
					// Greating
					s.ChannelMessageSend(m.ChannelID, "Hi "+m.Author.Username)
				}
			}
		}

		if strings.HasPrefix(m.Content, "!create") {
			SERVICE.CreateItem(s, m)
		} else if strings.HasPrefix(m.Content, "!set") {
			SERVICE.SetSerie(s, m)
		}

	}
}
