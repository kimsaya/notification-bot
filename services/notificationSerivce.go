package services

import (
	"fmt"
	"strings"

	HELPER "notification-bot/helpers"
	MODEL "notification-bot/models"
	REPO "notification-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func Users(s *discordgo.Session, m *discordgo.MessageCreate) {
	if user := REPO.FindUserByID(m.Author.ID); user != nil {
		user.LastActive = HELPER.GetNowTimestamp()
		REPO.UpdateUser(user)
		return
	}
	REPO.CreateUser(&MODEL.User{
		ID:         m.Author.ID,
		Name:       m.Author.Username,
		JointDate:  HELPER.GetNowTimestamp(),
		LastActive: HELPER.GetNowTimestamp(),
	})
}

// CreateItem a
func CreateItem(s *discordgo.Session, m *discordgo.MessageCreate) {
	messageSegment := strings.Split(m.Content, "\"")
	fmt.Print(messageSegment)
	// s.GuildRoles()
	role, err := s.GuildRoleCreate(m.GuildID)
	if err != nil {
		fmt.Println(err)
		return
	}
	role, err = s.GuildRoleEdit(m.GuildID, role.ID, messageSegment[1], 4171230, role.Hoist, role.Permissions, role.Mentionable)
	if err != nil {
		fmt.Println(err)
		return
	}
	var messageSend = new(discordgo.MessageSend)
	messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
	messageSend.AllowedMentions.Roles = append(messageSend.AllowedMentions.Roles, role.ID)
	messageSend.Content = "Created: <@&" + role.ID + ">"
	_, err = s.ChannelMessageSendComplex(m.ChannelID, messageSend)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// SetSerie a
func SetSerie(s *discordgo.Session, m *discordgo.MessageCreate) {
	// !set @Demon Slayer [7d]
	messageSegment := strings.FieldsFunc(m.Content, OptionSpliter)
	if len(messageSegment) < 2 {
		var messageSend = new(discordgo.MessageSend)
		messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
		messageSend.AllowedMentions.Users = append(messageSend.AllowedMentions.Users, m.Author.ID)
		messageSend.Content = "Dear <@" + m.Author.ID + ">: Your command is wrong. \n\tTry to set duration like this => [7d]"
		SendMessage(s, m, messageSend)
		return
	}
	if len(m.MentionRoles) < 1 {
		var messageSend = new(discordgo.MessageSend)
		messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
		messageSend.AllowedMentions.Users = append(messageSend.AllowedMentions.Users, m.Author.ID)
		messageSend.Content = "Dear <@" + m.Author.ID + ">: \n\tPls mention the item."
		SendMessage(s, m, messageSend)
		return
	}
	role := m.MentionRoles[0]

	// Check
	if REPO.FindItemByID(role) != nil {
		var messageSend = new(discordgo.MessageSend)
		messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
		messageSend.AllowedMentions.Users = append(messageSend.AllowedMentions.Users, m.Author.ID)
		messageSend.AllowedMentions.Roles = append(messageSend.AllowedMentions.Roles, role)
		messageSend.Content = "Dear <@" + m.Author.ID + ">: The process Set <@&" + role + ">" + " is Feild. Becuase: `Already set`."
		SendMessage(s, m, messageSend)
		return
	}

	item := new(MODEL.Item)
	item.ID = role
	item.Status = "Translation"
	item.Duration = GetDuration(messageSegment[1])
	if err := REPO.CreateItem(item); err != nil {
		fmt.Println(err)
		return
	}
	var messageSend = new(discordgo.MessageSend)
	messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
	messageSend.AllowedMentions.Roles = append(messageSend.AllowedMentions.Users, m.Author.ID)
	messageSend.AllowedMentions.Roles = append(messageSend.AllowedMentions.Roles, item.ID)
	messageSend.Content = ">>> Set: <@&" + item.ID + ">" + " Duration:" + HELPER.GetDurationFromTimestap(item.Duration) + " is Done. <@" + m.Author.ID + ">"
	SendMessage(s, m, messageSend)
}

//
func CheckSerie(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(m.MentionRoles) < 1 {
		var messageSend = new(discordgo.MessageSend)
		messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
		messageSend.AllowedMentions.Users = append(messageSend.AllowedMentions.Users, m.Author.ID)
		messageSend.Content = "Dear <@" + m.Author.ID + ">: \n\tPls mention the item."
		SendMessage(s, m, messageSend)
		return
	}
	role := m.MentionRoles[0]
	if item := REPO.FindItemByID(role); item != nil {
		var messageSend = new(discordgo.MessageSend)
		messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
		messageSend.AllowedMentions.Users = append(messageSend.AllowedMentions.Users, m.Author.ID)
		messageSend.AllowedMentions.Roles = append(messageSend.AllowedMentions.Roles, role)
		messageSend.Content = "Dear <@" + m.Author.ID +
			">\n>>> The <@&" + role + "> " +
			"\nTime Left: " + HELPER.GetDurationFromTimestap((item.CreatedDate+item.Duration)-HELPER.GetNowTimestamp()) +
			"\nStatus:    " + item.Status + "[WILLBETABLE]" +
			"\nLog:\n" + item.Description +
			"\nAssign :   Unknow For Now"
		SendMessage(s, m, messageSend)
		return
	} else {
		var messageSend = new(discordgo.MessageSend)
		messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
		messageSend.AllowedMentions.Users = append(messageSend.AllowedMentions.Users, m.Author.ID)
		messageSend.AllowedMentions.Roles = append(messageSend.AllowedMentions.Roles, role)
		messageSend.Content = "Dear <@" + m.Author.ID + ">\nThe <@&" + role + "> Never been `set` before."
		SendMessage(s, m, messageSend)
		return
	}
}

func SendMessage(s *discordgo.Session, m *discordgo.MessageCreate, ms *discordgo.MessageSend) {
	if _, err := s.ChannelMessageSendComplex(m.ChannelID, ms); err != nil {
		fmt.Println(err)
		return
	}
}

func OptionSpliter(r rune) bool {
	return r == '[' || r == ']'
}
func GetDuration(input string) int64 {
	var result int64 = 0
	if strings.Contains(input, "w") {
		input = strings.ReplaceAll(input, "w", "")
		result = HELPER.StringToInt64(input) * HELPER.GetOneWeek()
	} else if strings.Contains(input, "d") {
		input = strings.ReplaceAll(input, "d", "")
		result = HELPER.StringToInt64(input) * HELPER.GetOneDay()
	}
	return result
}
