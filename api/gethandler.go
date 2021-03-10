package api

import (
	HELPER "notification-bot/helpers"
	MODEL "notification-bot/models"
	REPO "notification-bot/repository"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if HELPER.StringContain(botChannel, m.ChannelID) {
		if strings.HasPrefix(m.Content, "!set") {
			//!Setup the Task
			user := REPO.FindUserByID(m.Author.ID)
			if user == nil {
				REPO.CreateUser(MODEL.User{Id: m.Author.ID, Name: m.Author.Username})
				s.ChannelMessageSend(m.ChannelID, "Welcome "+m.Author.Username)
			} else {
				s.ChannelMessageSend(m.ChannelID, "Hi "+user.Name)
			}
		}
	}
}
