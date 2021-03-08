package utils

import (
	"io"
	"log"

	HELPER "notification-bot/helpers"
)

var logsDirectory = ""

// InitalizeLog inital the log process.
func InitalizeLog(directory string) bool {
	logsDirectory = directory
	filePath := directory + "/" + HELPER.GetNowDate()
	writer := io.Writer(OpenFile(filePath))
	log.SetOutput(writer)
	log.Println("[IN] _UTILS initialized")
	return true
}

// checkOutDateLogFile find out of date log then remove to save storage
func checkOutDateLogFile() bool {

	return false
}
