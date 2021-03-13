package repository

import (
	"errors"
	HELPER "notification-bot/helpers"
	MODEL "notification-bot/models"
	UTILS "notification-bot/utils"
	"strings"
)

var userPath = "users/"

// CreateUser a
func CreateUser(user *MODEL.User) error {
	return saveUser(user)
}

//
func UpdateUser(user *MODEL.User) error {
	return saveUser(user)
}

// FindUserByID a
func FindUserByID(id string) *MODEL.User {
	for _, filePath := range UTILS.GetNoLimitInDirectory(repoDirectory + userPath) {
		userID := UTILS.GetFileNameFromPath(filePath)
		if userID == id {
			return readUser(filePath)
		}
	}
	return nil
}

func saveUser(user *MODEL.User) error {
	filePath := repoDirectory + userPath + user.ID
	value := "" +
		user.Name +
		"|" +
		HELPER.Int64ToString(user.LastActive) +
		"|" +
		HELPER.Int64ToString(user.JointDate) +
		""
	file := UTILS.OpenFile(filePath)
	file.Close()
	if UTILS.WriteFile(filePath, value) {
		return nil
	}
	return errors.New("Save User Failed.")
}
func readUser(filePath string) *MODEL.User {
	if content, status := UTILS.ReadFile(filePath); status == true {
		segments := strings.Split(content, "|")
		if len(segments) < 3 {
			return nil
		}
		return &MODEL.User{
			ID:         UTILS.GetFileNameFromPath(filePath),
			Name:       segments[0],
			LastActive: HELPER.StringToInt64(segments[1]),
			JointDate:  HELPER.StringToInt64(segments[2]),
		}
	}
	return nil
}
