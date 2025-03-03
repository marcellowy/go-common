package tools

import (
	"fmt"
	"time"
)

const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
	PB = TB * 1024
)

func FormatBytes(bytes uint64) string {
	switch {
	case bytes >= TB:
		return fmt.Sprintf("%.2f TB", float64(bytes)/float64(TB))
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}

func FormatDuration(d time.Duration) string {
	days := d / (24 * time.Hour)
	d %= 24 * time.Hour
	hours := d / time.Hour
	d %= time.Hour
	minutes := d / time.Minute
	d %= time.Minute
	seconds := d / time.Second

	result := ""
	s := ""
	if days > 0 {
		result += fmt.Sprintf("%d day", days)
	}
	if hours > 0 {
		s = "%d hour"
		if len(result) > 0 {
			s = " %d hour"
		}
		result += fmt.Sprintf(s, hours)
	}
	if minutes > 0 {
		s = "%d min"
		if len(result) > 0 {
			s = " %d min"
		}
		result += fmt.Sprintf(s, minutes)
	}
	if seconds > 0 || result == "" { // 如果时间太短，确保至少显示 "0秒"
		s = "%d s"
		if len(result) > 0 {
			s = " %d s"
		}
		result += fmt.Sprintf(s+
			"", seconds)
	}
	return result
}
