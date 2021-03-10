package utils

import (
	"io"
	"log"

	HELPER "notification-bot/helpers"
)

// initalizeLog inital the log process.
func initalizeLog(logDirectory string) bool {

	filePath := logDirectory + "/" + HELPER.GetNowDate()
	writer := io.Writer(OpenFile(filePath))
	log.SetOutput(writer)
	return true
}

// checkOutDateLogFile find out of date log then remove to save storage\
// duration is how long the file will keep in day.
func checkOutDateLogFile(logsDirectory string, duration int) bool {

	filePaths := GetNoLimitInDirectory(logsDirectory)
	for _, filePath := range filePaths {
		if HELPER.GetTimestampFromStringOfDate(GetFileNameFromPath(filePath)) < HELPER.GetNowTimestamp()-(int64(duration)*HELPER.GetOneDay()) {
			RemoveFile(filePath)
		}
	}
	return false
}
