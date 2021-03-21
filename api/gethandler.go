package api

import (
	"fmt"
	HELPER "notification-bot/helpers"
	SERVICE "notification-bot/services"

	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	// Check Channel
	if HELPER.StringContains(botChannel, m.ChannelID) {
		// If Bot it self end progres
		if m.Author.ID == s.State.User.ID {
			return
		}
		if strings.Contains(m.Content, "-_-") {
			s.ChannelMessageSend(m.ChannelID, "I'm sorry.\nI try my best.")
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
		SERVICE.Users(s, m)
		if strings.HasPrefix(m.Content, "!create") {
			SERVICE.CreateItem(s, m)
		} else if strings.HasPrefix(m.Content, "!set") {
			SERVICE.SetSerie(s, m)
		} else if strings.HasPrefix(m.Content, "!check") {
			SERVICE.CheckSerie(s, m)
		} else if strings.HasPrefix(m.Content, "!translate") ||
			strings.HasPrefix(m.Content, "!edit") ||
			strings.HasPrefix(m.Content, "!post") {
			SERVICE.CreateTask(s, m)
		}

	}
}

func messageReaction(s *discordgo.Session, m *discordgo.MessageReactionAdd) {

	fmt.Println(m.Emoji)
	// "❤️"
	//"✅ ❌"

}
