package models

import "strings"

// Task a
type Task struct {
	ID          string `json:"id"` //UserID-BookID-Index-Activity
	Name        string `json:"name"`
	UserID      string `json:"user_id"`
	ItemID      string `json:"item_id"`
	Status      string `json:"status"`
	LastUpdate  int64  `json:"last_update"`
	CreatedDate int64  `json:"creater_date"`
}

type TaskDTO struct {
	Name         string
	IsTranslated bool
	IsEdited     bool
	IsPosted     bool
}

func TaskToTaskDTO(input []Task) []TaskDTO {
	var taskDTOs []TaskDTO
	for _, task := range input {
		tempTaskDTO := new(TaskDTO)
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
	return taskDTOs
}

func TaskDTOsToString(input []TaskDTO) string {
	var value = ""
	for _, taskDTO := range input {
		value += "\n"
		value += " Translated"
		if taskDTO.IsTranslated {
			value += "[✅]"
		} else {
			value += "[❌]"
		}
		value += " Edit"
		if taskDTO.IsEdited {
			value += "[✅]"
		} else {
			value += "[❌]"
		}
		value += " Post"
		if taskDTO.IsPosted {
			value += "[✅]"
		} else {
			value += "[❌]"
		}
		value += " | " + taskDTO.Name
	}
	return value
}
