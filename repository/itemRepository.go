package repository

import (
	"errors"
	"log"
	HELPER "notification-bot/helpers"
	MODEL "notification-bot/models"
	UTILS "notification-bot/utils"
	"strings"
)

var itemPath = "items/"

func CreateItem(item *MODEL.Item) error {
	item.CreatedDate = HELPER.GetNowTimestamp()
	return saveItem(item)
}

func UpdateItem(item *MODEL.Item) error {
	item.LastUpdate = HELPER.GetNowTimestamp()
	return saveItem(item)
}

func FindItemByID(id string) (*MODEL.Item, error) {
	for _, filePath := range UTILS.GetFileNoLimitInDirectory(repoDirectory + itemPath) {
		itemID := UTILS.GetFileNameFromPath(filePath)
		if itemID == id {
			return readItem(itemID)
		}
	}
	return nil, nil
}

func readItem(id string) (*MODEL.Item, error) {

	filePath := repoDirectory + itemPath + id
	Content, status := UTILS.ReadFile(filePath)
	if status {
		item := new(MODEL.Item)
		segments := strings.Split(Content, "|")
		if len(segments) < 7 {
			log.Println("[ER] Read Item: Bad File")
			return nil, errors.New("Bad File")
		}
		item.ID = id
		item.Name = segments[0]
		item.TranslatorID = segments[1]
		item.EditorID = segments[2]
		item.PostorID = segments[3]
		item.Duration = HELPER.StringToInt64(segments[4])
		item.LastUpdate = HELPER.StringToInt64(segments[5])
		item.CreatedDate = HELPER.StringToInt64(segments[6])
		return item, nil
	}
	return nil, nil
}

func saveItem(item *MODEL.Item) error {

	filePath := repoDirectory + itemPath + item.ID
	file := UTILS.OpenFile(filePath)
	file.Close()
	value := item.Name +
		"|" +
		item.TranslatorID +
		"|" +
		item.EditorID +
		"|" +
		item.PostorID +
		"|" +
		HELPER.Int64ToString(item.Duration) +
		"|" +
		HELPER.Int64ToString(item.LastUpdate) +
		"|" +
		HELPER.Int64ToString(item.CreatedDate) +
		""
	if UTILS.WriteFile(filePath, value) {
		return nil
	}
	log.Println("[ER] Save Item: Failed", item.ID)
	return errors.New("Can't save Item: " + filePath)
}

func LogItem(id, message string) {
	filePath := repoDirectory + itemPath + id + "-LOG/" + HELPER.GetNowDate()
	file := UTILS.OpenFile(filePath)
	file.Close()
	message = HELPER.GetNowTime() + " :" + message
	UTILS.AppendFile(filePath, message)
}
