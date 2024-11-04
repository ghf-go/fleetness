package utils

import "time"

const (
	T_DATE     = "2006-01-02"
	T_DATE_HM  = "2006-01-02 15:04"
	T_TIME     = "15:04:05.999"
	T_DATETIME = "2006-01-02 15:04:05.999"
)

func FormatDateTime(t ...time.Time) string {
	if len(t) > 0 {
		return t[0].Format(T_DATETIME)
	}
	return time.Now().Format(T_DATETIME)
}
