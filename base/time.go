package base

import "time"

const TimeFormatString = "2006-01-02 15:04:05"

// FormatTime 将时间缀格式为YYYY-MM-DD HH:mm:ss
func FormatTime(i int64) string {
	return time.Unix(i, 0).Format(TimeFormatString)
}
