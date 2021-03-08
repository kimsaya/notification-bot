package helpers

import "time"

// GetNow return the current datetime as Go.Time
func GetNow() time.Time {
	return time.Now()
}

// GetNowDate return current date as string with format yyyy-MM-dd
func GetNowDate() (reDate string) {
	return GetNow().Format("2006-01-02") //2006-01-02 is a format yyyy-MM-dd
}

// GetNowTime return current time as string with format HH:mm:ss
func GetNowTime() (reTime string) {
	return GetNow().Format("15:04:05") //15:04:05 is a format HH:mm:ss
}

// GetNowTimestamp return current datetime as timestamp
func GetNowTimestamp() (reTimeStamp int64) {
	return GetNow().UnixNano() / int64(time.Millisecond)
}
