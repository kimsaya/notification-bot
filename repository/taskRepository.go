package repository

import (
	"errors"
	HELPER "notification-bot/helpers"
	MODEL "notification-bot/models"
	UTILS "notification-bot/utils"
	"strings"
)

var taskPath = "tasks/"

func CreateTask(task *MODEL.Task) error {

	task.CreatedDate = HELPER.GetNowTimestamp()
	task.LastUpdate = task.CreatedDate
	if task.ID == "" || len(strings.Split(task.ID, "-")) < 4 {
		return errors.New("Invalid ID")
	}
	return saveTask(task)
}

func UpdateTask(task *MODEL.Task) error {
	task.LastUpdate = HELPER.GetNowTimestamp()
	if task.ID == "" || len(strings.Split(task.ID, "-")) < 4 {
		return errors.New("Invalid ID")
	}
	return saveTask(task)
}

func FindTaskByID(id string) (*MODEL.Task, error) {
	segments := strings.Split(id, "-")
	filePath := repoDirectory + taskPath + segments[0] + "-" + segments[1] + "/" + id
	return readTask(filePath)
}

func FindTaskByUserID(userID string) (*[]MODEL.Task, error) {
	var tasks []MODEL.Task
	for _, dirPath := range UTILS.GetDirectoryNoLimitInDirectory(repoDirectory + taskPath) {
		userIDItemID := UTILS.GetFileNameFromPath(dirPath)
		segments := strings.Split(userIDItemID, "-")
		if len(segments) < 2 {
			return nil, errors.New("Invalid ID")
		}
		if segments[0] != userID {
			continue
		}
		for _, filePath := range UTILS.GetFileNoLimitInDirectory(dirPath) {
			task, _ := readTask(filePath)
			tasks = append(tasks, *task)
		}

	}
	return &tasks, nil
}

func FindTaskByItemID(itemID string) (*[]MODEL.Task, error) {
	var tasks []MODEL.Task
	for _, dirPath := range UTILS.GetDirectoryNoLimitInDirectory(repoDirectory + taskPath) {
		userIDItemID := UTILS.GetFileNameFromPath(dirPath)
		segments := strings.Split(userIDItemID, "-")
		if len(segments) < 2 {
			return nil, errors.New("Invalid ID")
		}
		if segments[1] != itemID {
			continue
		}
		for _, filePath := range UTILS.GetFileNoLimitInDirectory(dirPath) {
			task, _ := readTask(filePath)
			tasks = append(tasks, *task)
		}

	}
	return &tasks, nil
}

func saveTask(task *MODEL.Task) error {
	task.ID = strings.ReplaceAll(task.ID, " ", "#")
	filePath := repoDirectory + taskPath + task.UserID + "-" + task.ItemID + "/" + task.ID
	value := "" +
		task.Status +
		"|" +
		task.Name +
		"|" +
		HELPER.Int64ToString(task.LastUpdate) +
		"|" +
		HELPER.Int64ToString(task.CreatedDate) +
		""
	file := UTILS.OpenFile(filePath)
	file.Close()
	if UTILS.WriteFile(filePath, value) {
		return nil
	}
	return errors.New("Save Task Failed.")
}

func readTask(filePath string) (*MODEL.Task, error) {
	if content, status := UTILS.ReadFile(filePath); status == true {
		segments := strings.Split(content, "|")
		if len(segments) < 4 {
			return nil, errors.New("Bad File")
		}
		tempID := strings.ReplaceAll(UTILS.GetFileNameFromPath(filePath), "#", " ")
		return &MODEL.Task{
			ID:          tempID,
			Status:      segments[0],
			Name:        segments[1],
			LastUpdate:  HELPER.StringToInt64(segments[2]),
			CreatedDate: HELPER.StringToInt64(segments[3]),
		}, nil
	}
	return nil, nil
}
