package helpers

import (
	"log"
	"time"
)

// GetNow return the current datetime as Go.Time
func GetNow() time.Time {
	return time.Now()
}

// GetNowDate return current date as string with format yyyy-MM-dd
func GetNowDate() string {
	return GetNow().Format("2006-01-02") //2006-01-02 is a format yyyy-MM-dd
}

// GetNowTime return current time as string with format HH:mm:ss
func GetNowTime() string {
	return GetNow().Format("15:04:05") //15:04:05 is a format HH:mm:ss
}

// GetNowTimestamp return current datetime as timestamp
func GetNowTimestamp() int64 {
	return GetNow().UnixNano() / int64(time.Millisecond)
}

// GetTimestampFromTime turn Time object in to Timestamp
func GetTimestampFromTime(inTime time.Time) int64 {
	return inTime.UnixNano() / int64(time.Millisecond)
}

// GetTimestampFromStringOfDate turn input string as yyyy-MM-dd format to Timestamp
func GetTimestampFromStringOfDate(input string) int64 {
	layout := "2006-01-02"
	t, err := time.Parse(layout, input)
	if err != nil {
		log.Println(err)
		//fmt.Println(err)
		return 0
	}
	return t.UnixNano() / int64(time.Millisecond)
}

// GetOneDay return Timestamp in one day
func GetOneDay() int64 {
	return 86400000
}
