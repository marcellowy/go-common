// Package vtime
// Copyright 2023 marcello<volibearw@gmail.com>. All rights reserved.
package vtime

import (
	"golang.org/x/exp/constraints"
	"time"
)

// TimeFormatString time layout
const TimeFormatString = "2006-01-02 15:04:05"

// FormatTime format time use TimeFormatString layout
func FormatTime[T constraints.Float | constraints.Integer](sec T, nsec T) string {
	return time.Unix(int64(sec), int64(nsec)).Format(TimeFormatString)
}

// ParseInAsiaShanghai parse time
func ParseInAsiaShanghai(s, formatString string) (t time.Time, err error) {
	var (
		locShanghai *time.Location
	)
	if locShanghai, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		return
	}
	if t, err = time.ParseInLocation(formatString, s, locShanghai); err != nil {
		return
	}
	return
}
