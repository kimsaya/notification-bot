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
	item.Description = "- Creating: " + HELPER.GetNowDate() + " " + HELPER.GetNowTime() + "\n"
	item.CreatedDate = HELPER.GetNowTimestamp()
	return saveItem("", item)
}

func UpdateItem(item *MODEL.Item) error {
	item.Description += "- Update: " + HELPER.GetNowDate() + " " + HELPER.GetNowTime() + "\n"
	item.LastUpdate = HELPER.GetNowTimestamp()
	return saveItem("", item)
}

func FindItemByID(id string) *MODEL.Item {
	for _, filePath := range UTILS.GetNoLimitInDirectory(repoDirectory + itemPath) {
		itemID := UTILS.GetFileNameFromPath(filePath)
		if itemID == id {
			return readItem("", itemID)
		}
	}
	return nil
}

func readItem(mainID, id string) *MODEL.Item {
	if mainID != "" {
		// SubItem
		mainID += "-sub/"
	}
	filePath := repoDirectory + itemPath + mainID + id
	Content, status := UTILS.ReadFile(filePath)
	if status {
		item := new(MODEL.Item)
		segments := strings.Split(Content, "|")
		if len(segments) < 6 {
			log.Println("")
			return nil
		}
		item.ID = id
		item.Name = segments[0]
		item.Description = segments[1]
		item.Status = segments[2]
		item.SubItems = new([]MODEL.Item)
		if mainID == "" {
			for _, subFilePath := range UTILS.GetNoLimitInDirectory(repoDirectory + itemPath + id + "-sub") {
				*item.SubItems = append(*item.SubItems, *readItem(id, UTILS.GetFileNameFromPath(subFilePath)))
			}
		}
		item.Duration = HELPER.StringToInt64(segments[3])
		item.LastUpdate = HELPER.StringToInt64(segments[4])
		item.CreatedDate = HELPER.StringToInt64(segments[5])
		return item
	}
	return nil
}

func saveItem(id string, item *MODEL.Item) error {
	if id != "" {
		// SubItem
		id += "-sub/"
	}
	filePath := repoDirectory + itemPath + id + item.ID
	file := UTILS.OpenFile(filePath)
	file.Close()

	if item.SubItems != nil {
		for _, subItem := range *item.SubItems {
			saveItem(item.ID, &subItem)
		}
	}
	value := item.Name +
		"|" +
		item.Description +
		"|" +
		item.Status +
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
