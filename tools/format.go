package tools

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
	PB = TB * 1024
)

// ParseSizeUnit
// parse 10M/10MB/10m/10mb to byte or other kb/mb/gb/tb
func ParseSizeUnit(s string) (byte int64, err error) {
	var (
		size int64
		unit string
	)
	s = strings.TrimSpace(s)
	var reg *regexp.Regexp
	if reg, err = regexp.Compile("([0-9]+)([A-Za-z]+)"); err != nil {
		return
	}
	match := reg.FindStringSubmatch(s)
	if len(match) != 3 {
		err = fmt.Errorf("parse size unit error")
		return
	}
	if size, err = strconv.ParseInt(match[1], 10, 64); err != nil {
		return
	}
	unit = match[2]
	if len(unit) == 1 {
		unit += "b"
	}

	switch strings.ToLower(unit) {
	case "kb":
		size = size * KB
	case "mb":
		size = size * MB
	case "gb":
		size = size * GB
	case "tb":
		size = size * TB
	default:
		return 0, fmt.Errorf("parse unit error")
	}
	return size, nil
}

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
