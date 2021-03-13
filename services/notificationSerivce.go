package services

import (
	"fmt"
	"strings"

	HELPER "notification-bot/helpers"
	MODEL "notification-bot/models"

	"github.com/bwmarrin/discordgo"
)

// CreateItem a
func CreateItem(s *discordgo.Session, m *discordgo.MessageCreate) {
	messageSegment := strings.Split(m.Content, "\"")
	fmt.Print(messageSegment)
	// s.GuildRoles()
	rs, err := s.GuildRoleCreate(m.GuildID)
	if err != nil {
		fmt.Println(err)
		return
	}
	rs, err = s.GuildRoleEdit(m.GuildID, rs.ID, messageSegment[1], 4171230, rs.Hoist, rs.Permissions, rs.Mentionable)
	if err != nil {
		fmt.Println(err)
		return
	}
	var messageSend = new(discordgo.MessageSend)
	messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
	messageSend.AllowedMentions.Roles = append(messageSend.AllowedMentions.Roles, rs.ID)
	messageSend.Content = "Created: <@&" + rs.ID + ">"
	_, err = s.ChannelMessageSendComplex(m.ChannelID, messageSend)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// SetSerie a
func SetSerie(s *discordgo.Session, m *discordgo.MessageCreate) {
	// !set @Demon Slayer [7d]
	// messageSegment := strings.Split(m.Content, "\"")
	role := m.MentionRoles[0]
	item := new(MODEL.Item)
	item.ID = role
	item.Status = "Translation"
	item.CreatedDate = HELPER.GetNowTimestamp()
}
