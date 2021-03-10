package repository

import (
	MODEL "notification-bot/models"
	UTILS "notification-bot/utils"
)

// CreateUser a
func CreateUser(user MODEL.User) bool {
	filePath := repoDirectory + "/users/" + user.Id
	file := UTILS.OpenFile(filePath)
	file.Close()
	return UTILS.WriteFile(filePath, user.Name)
}

// FindUserByID a
func FindUserByID(id string) *MODEL.User {
	for _, filePath := range UTILS.GetNoLimitInDirectory(repoDirectory + "/users/") {
		userID := UTILS.GetFileNameFromPath(filePath)
		if userID == id {
			name, status := UTILS.ReadFile(filePath)
			if status {
				return &MODEL.User{Id: userID, Name: name}
			}
		}
	}
	return nil
}
