package model

import "time"

const (
	FormatDate = "2006-01-02 15:04:05"
)

func UnixToDate(ts int64) string {
	return time.Unix(ts, 0).Format(FormatDate)
}

func Unix() int64 {
	return time.Now().Unix()
}
func Date() string {
	return time.Now().Format(FormatDate)
}
