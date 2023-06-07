// Package tools
// Copyright 2023 marcello<volibearw@gmail.com>. All rights reserved.
package tools

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
