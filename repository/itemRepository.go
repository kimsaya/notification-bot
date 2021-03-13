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

	filePath := repoDirectory + itemPath + item.ID
	file := UTILS.OpenFile(filePath)
	defer file.Close()

	item.Description = "- Creating: " + HELPER.GetNowDate() + " " + HELPER.GetNowTime() + "\n"
	item.CreatedDate = HELPER.GetNowTimestamp()
	return saveItem(filePath, item)
}

func FindItemByID(id string) *MODEL.Item {
	for _, filePath := range UTILS.GetNoLimitInDirectory(repoDirectory + itemPath) {
		itemID := UTILS.GetFileNameFromPath(filePath)
		if itemID == id {
			return readItem(filePath)
		}
	}
	return nil
}

func readItem(filePath string) *MODEL.Item {
	Content, status := UTILS.ReadFile(filePath)
	if status {
		item := new(MODEL.Item)
		segments := strings.Split(Content, "|")
		if len(segments) < 6 {
			log.Println("")
			return nil
		}
		item.Name = segments[0]
		item.Description = segments[1]
		item.Status = segments[2]
		subsegments := strings.Split(segments[3], ",")
		for _, subsegment := range subsegments {
			item.SubItems = append(item.SubItems, subsegment)
		}
		item.Duration = HELPER.StringToInt64(segments[4])
		item.LastUpdate = HELPER.StringToInt64(segments[5])
		item.CreatedDate = HELPER.StringToInt64(segments[6])
		return item
	}
	return nil
}
func saveItem(filePath string, item *MODEL.Item) error {
	sub := ""
	for _, loop := range item.SubItems {
		sub += loop + ","
	}
	value := item.Name +
		"|" +
		item.Description +
		"|" +
		item.Status +
		"|" +
		sub +
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
	return errors.New("Can't save Item: " + filePath)
}
