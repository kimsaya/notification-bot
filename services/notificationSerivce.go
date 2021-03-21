package services

import (
	"fmt"
	"log"
	"sort"
	"strings"

	HELPER "notification-bot/helpers"
	MODEL "notification-bot/models"
	REPO "notification-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func Greating(s *discordgo.Session, channerlID string) {

	return

	// users := REPO.FindAllUsers()

	// var messageSend = new(discordgo.MessageSend)
	// messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
	// messageSend.Content = HELPER.GetNowTimeName() + " \n>>> "
	// for _, user := range *users {
	// 	messageSend.AllowedMentions.Users = append(messageSend.AllowedMentions.Users, user.ID)
	// 	messageSend.Content += "<@" + user.ID + "> \n"
	// }
	// _, err := s.ChannelMessageSendComplex(channerlID, messageSend)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
}

func Users(s *discordgo.Session, m *discordgo.MessageCreate) {
	if user, _ := REPO.FindUserByID(m.Author.ID); user != nil {
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
	// messageSegment := strings.Split(m.Content, "\"")
	// role, err := s.GuildRoleCreate(m.GuildID)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// role, err = s.GuildRoleEdit(m.GuildID, role.ID, messageSegment[1], 4171230, role.Hoist, role.Permissions, role.Mentionable)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// var messageSend = new(discordgo.MessageSend)
	// messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
	// messageSend.AllowedMentions.Roles = append(messageSend.AllowedMentions.Roles, role.ID)
	// messageSend.Content = "Created: <@&" + role.ID + ">"
	// _, err = s.ChannelMessageSendComplex(m.ChannelID, messageSend)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
}

// SetSerie is Create a Seria
func SetSerie(s *discordgo.Session, m *discordgo.MessageCreate) {
	// !set @Demon Slayer [7d]
	messageSegment := strings.FieldsFunc(m.Content, OptionSpliter)

	if len(messageSegment) < 2 {
		var messageSend = new(discordgo.MessageSend)
		messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
		messageSend.AllowedMentions.Users = append(messageSend.AllowedMentions.Users, m.Author.ID)
		messageSend.Content = "Dear <@" + m.Author.ID + ">: Your command is wrong. " +
			"\n\tExample: !set @Series [7d] " +
			"\n\t[1d] or [1] = 1 day"
		SendMessage(s, m, messageSend)
		return
	}
	duration := GetDuration(messageSegment[1])
	if duration <= 0 {
		var messageSend = new(discordgo.MessageSend)
		messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
		messageSend.AllowedMentions.Users = append(messageSend.AllowedMentions.Users, m.Author.ID)
		messageSend.Content = "Dear <@" + m.Author.ID + ">: Your command is wrong. " +
			"\n\tExample: !set @Series [7d] " +
			"\n\t[1d] or [1] = 1 day"
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
	role, _ := s.State.Role(m.GuildID, m.MentionRoles[0])

	// Check
	if item, _ := REPO.FindItemByID(role.ID); item != nil {
		var messageSend = new(discordgo.MessageSend)
		messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
		messageSend.AllowedMentions.Users = append(messageSend.AllowedMentions.Users, m.Author.ID)
		messageSend.AllowedMentions.Roles = append(messageSend.AllowedMentions.Roles, role.ID)
		messageSend.Content = "Dear <@" + m.Author.ID + ">: The <@&" + role.ID + ">" + " is `Already set`." +
			"\n\t Duration: " + HELPER.GetDurationFromTimestap(item.Duration)
		SendMessage(s, m, messageSend)
		return
	}

	item := new(MODEL.Item)
	item.ID = role.ID
	item.Name = role.Name
	item.Duration = duration
	if err := REPO.CreateItem(item); err != nil {
		log.Println(err)
		var messageSend = new(discordgo.MessageSend)
		messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
		messageSend.AllowedMentions.Users = append(messageSend.AllowedMentions.Users, m.Author.ID)
		messageSend.AllowedMentions.Roles = append(messageSend.AllowedMentions.Roles, role.ID)
		messageSend.Content = "Dear <@" + m.Author.ID + ">: The process failed." +
			"\n\t`" + err.Error() + "`"
		SendMessage(s, m, messageSend)
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
func CreateTask(s *discordgo.Session, m *discordgo.MessageCreate) {
	// !translate @series [Chapter X]
	messageSegment := strings.FieldsFunc(m.Content, OptionSpliter)
	if len(messageSegment) < 2 {
		var messageSend = new(discordgo.MessageSend)
		messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
		messageSend.AllowedMentions.Users = append(messageSend.AllowedMentions.Users, m.Author.ID)
		messageSend.Content = "Dear <@" + m.Author.ID + ">: Your command is wrong." +
			"\n\tExample: !translate @Series [Chapter X]"
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
	item, _ := REPO.FindItemByID(role)
	if item == nil {
		var messageSend = new(discordgo.MessageSend)
		messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
		messageSend.AllowedMentions.Users = append(messageSend.AllowedMentions.Users, m.Author.ID)
		messageSend.AllowedMentions.Roles = append(messageSend.AllowedMentions.Roles, role)
		messageSend.Content = "Dear <@" + m.Author.ID + ">: The <@&" + role + ">" + " is `Not been set yet`."
		SendMessage(s, m, messageSend)
		return
	}

	// Find Exiting

	newTask := new(MODEL.Task)
	newTask.UserID = m.Author.ID
	newTask.ItemID = item.ID
	newTask.Name = messageSegment[1]
	newTask.ID = newTask.UserID + "-" + newTask.ItemID + "-" + messageSegment[1]

	if strings.HasPrefix(m.Content, "!translate") {
		newTask.ID += "-T"
		newTask.Status = "Translated"
	} else if strings.HasPrefix(m.Content, "!edit") {
		newTask.ID += "-E"
		newTask.Status = "Edited"
	} else if strings.HasPrefix(m.Content, "!post") {
		newTask.ID += "-P"
		newTask.Status = "Posted"
	}

	if err := REPO.CreateTask(newTask); err != nil {
		var messageSend = new(discordgo.MessageSend)
		messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
		messageSend.AllowedMentions.Users = append(messageSend.AllowedMentions.Users, m.Author.ID)
		messageSend.AllowedMentions.Roles = append(messageSend.AllowedMentions.Roles, role)
		messageSend.Content = "Dear <@" + m.Author.ID + ">: The process failed." +
			"\n\t`" + err.Error() + "`"
		SendMessage(s, m, messageSend)
		return
	}
	s.MessageReactionAdd(m.ChannelID, m.ID, "❤️")

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
	var user *discordgo.User
	if len(m.Mentions) > 0 {
		user = m.Mentions[0]
	}

	if item, _ := REPO.FindItemByID(role); item != nil {
		var messageSend = new(discordgo.MessageSend)
		messageSend.AllowedMentions = new(discordgo.MessageAllowedMentions)
		messageSend.AllowedMentions.Users = append(messageSend.AllowedMentions.Users, m.Author.ID)
		messageSend.AllowedMentions.Roles = append(messageSend.AllowedMentions.Roles, role)
		messageSend.Content = "Dear <@" + m.Author.ID +
			">\n>>> The <@&" + role + "> " +
			"\nTranslator \t: Unknow" +
			"\nEditor \t\t\t: Unknow" +
			"\nDuration\t\t: " + HELPER.GetDurationFromTimestap((item.CreatedDate+item.Duration)-HELPER.GetNowTimestamp()) +
			"\nStatus: "

		tasks := new([]MODEL.Task)
		if user == nil {
			tasks, _ = REPO.FindTaskByItemID(item.ID)

		}
		var taskDTOs []MODEL.TaskDTO
		for _, task := range *tasks {
			tempTaskDTO := new(MODEL.TaskDTO)
			for index, taskDTO := range taskDTOs {
				if taskDTO.Name == task.Name {
					tempTaskDTO = &taskDTOs[index]
					break
				}
			}
			if strings.HasSuffix(task.ID, "T") {
				tempTaskDTO.IsTranslated = true
			} else if strings.HasSuffix(task.ID, "E") {
				tempTaskDTO.IsEdited = true
			} else if strings.HasSuffix(task.ID, "P") {
				tempTaskDTO.IsPosted = true
			}
			if tempTaskDTO.Name == "" {
				tempTaskDTO.Name = task.Name
				taskDTOs = append(taskDTOs, *tempTaskDTO)
			}

		}

		sort.SliceStable(taskDTOs, func(i, j int) bool {
			isplit := strings.Split(taskDTOs[i].Name, " ")
			jsplit := strings.Split(taskDTOs[j].Name, " ")
			return HELPER.StringToInt64(isplit[1]) < HELPER.StringToInt64(jsplit[1])
		})

		for _, taskDTO := range taskDTOs {
			messageSend.Content += "\n"
			messageSend.Content += " Translated"
			if taskDTO.IsTranslated {
				messageSend.Content += "[✅]"
			} else {
				messageSend.Content += "[❌]"
			}
			messageSend.Content += " Edit"
			if taskDTO.IsEdited {
				messageSend.Content += "[✅]"
			} else {
				messageSend.Content += "[❌]"
			}
			messageSend.Content += " Post"
			if taskDTO.IsPosted {
				messageSend.Content += "[✅]"
			} else {
				messageSend.Content += "[❌]"
			}
			messageSend.Content += "  " + taskDTO.Name
		}
		// messageSend.Content += "\n\nLog:\n" + item.Description
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

// !assign @user [@series] translator/editor/poster
func AssignUserToSeries(s *discordgo.Session, m *discordgo.MessageCreate) {

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
	} else {
		result = HELPER.StringToInt64(input) * HELPER.GetOneDay()
	}
	return result
}
