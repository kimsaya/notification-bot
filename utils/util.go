package utils

// Initalize the util
func Initalize(logDirectory string, logDuration int) {
	initalizeLog(logDirectory)
	checkOutDateLogFile(logDirectory, logDuration)
}
