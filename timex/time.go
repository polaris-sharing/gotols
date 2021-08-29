package timex

import "time"

const dateFormat = "2006-01-02 15:04:05"

// TimeNow get the now time formate
func TimeNow() string {
	return time.Now().Format(dateFormat)
}
