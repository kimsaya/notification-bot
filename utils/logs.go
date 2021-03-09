package utils

import (
	"io"
	"log"

	HELPER "notification-bot/helpers"
)

// initalizeLog inital the log process.
func initalizeLog(logDirectory string) bool {

	filePath := logDirectory + "/" + HELPER.GetNowDate()
	writer := io.Writer(openFile(filePath))
	log.SetOutput(writer)
	log.Println("[IN] UTILS InitalizeLog")
	return true
}

// checkOutDateLogFile find out of date log then remove to save storage\
// duration is how long the file will keep in day.
func checkOutDateLogFile(logsDirectory string, duration int) bool {

	filePaths := getNoLimitInDirectory(logsDirectory)
	for _, filePath := range filePaths {
		if HELPER.GetTimestampFromStringOfDate(getFileNameFromPath(filePath)) < HELPER.GetNowTimestamp()-(int64(duration)*HELPER.GetOneDay()) {
			removeFile(filePath)
		}
	}
	return false
}
