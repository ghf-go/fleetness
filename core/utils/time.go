package utils

import (
	"time"
)

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

// 获取今天时间
func DayUnixMilli(ts ...time.Time) int64 {
	t := time.Now()
	if len(ts) > 0 {
		t = ts[0]

	}
	h, m, s := t.Clock()
	return t.UnixMilli() - int64((h*3600+m*60+s)*1000)

}

// 时间戳转时间
func UnixMilli2Time(m int64) time.Time {
	return time.Unix(m/1000, (m%1000)*1000000)
}

// 获取星期一的时间
func WeekStart(ts ...time.Time) time.Time {
	t := time.Now()
	if len(ts) > 0 {
		t = ts[0]
	}
	d := int(t.Weekday())
	ms := t.UnixNano()
	if d == 0 {
		ms -= (24*6 + int64(t.Hour())*int64(time.Hour)) + int64(t.Minute())*int64(time.Minute) + int64(t.Second())
	} else {
		ms -= (24*int64(d-1)+int64(t.Hour()))*int64(time.Hour) + int64(t.Minute())*int64(time.Minute) + int64(t.Second())
	}
	return UnixMilli2Time(ms / int64(time.Millisecond))

}

// 获取两个时间的周间隔
func SubWeeks(t1, t2 time.Time) uint {

	tt1 := WeekStart(t1)
	tt2 := WeekStart(t2)
	if tt2.Year() <= 1970 {
		return 1
	}
	w := tt1.Sub(tt2) / (time.Hour * time.Duration(24*7))
	// fmt.Printf("时间比较 %s -> %s  sub: %d , %d\n", tt1.Format(T_DATE), tt2.Format(T_DATE), tt1.Sub(tt2), w)
	return uint(w) + 1
}
