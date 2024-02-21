package utils

import "time"

func StringToTime(timeStr string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	return time.Parse(layout, timeStr)
}

func TimeToString(time time.Time) string {
	layout := "2006-01-02 15:04:05"
	return time.Format(layout)
}
