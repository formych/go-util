package timex

import (
	"time"
)

// 时间常量
const (
	Day   = 24 * time.Hour
	Week  = 7 * Day
	Month = 30 * Day
)

// GetCurrentYear 获取当前年份
func GetCurrentYear() int {
	return time.Now().Year()
}

// GetCurrentMonth 获取当前月份
func GetCurrentMonth() time.Month {
	return time.Now().Month()
}

// GetCurrentDay 获取当前日
func GetCurrentDay() int {
	return time.Now().Day()
}

// GetCurrentHour 获取当前小时
func GetCurrentHour() int {
	return time.Now().Hour()
}

// GetCurrentMinute 获取当前分钟
func GetCurrentMinute() int {
	return time.Now().Minute()
}

// GetCurrentSecond 获取当前分钟
func GetCurrentSecond() int {
	return time.Now().Second()
}

// GetCurrentWeekday 获取本周几
func GetCurrentWeekday() time.Weekday {
	return time.Now().Weekday()
}

// GetBeginningOfCurrentYear 获取当年的开始时间
func GetBeginningOfCurrentYear() time.Time {
	return time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.Local)
}

// GetBeginningOfCurrentMonth 获取当月的开始时间
func GetBeginningOfCurrentMonth() time.Time {
	year, month, _ := time.Now().Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
}

// GetBeginningOfCurrentWeek 获取本周的开始时间
func GetBeginningOfCurrentWeek() time.Time {
	year, month, day := time.Now().Date()
	weekday := int(time.Now().Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return time.Date(year, month, day-weekday+1, 0, 0, 0, 0, time.Local)
}

// GetBeginningOfYear 获取一年的开始时间
func GetBeginningOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// GetBeginningOfMonth 获取一个月的开始时间
func GetBeginningOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

// GetBeginningOfDay 获取一天的开始时间
func GetBeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// GetBeginningOfNextDay 获取一天的开始时间
func GetBeginningOfNextDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day+1, 0, 0, 0, 0, t.Location())
}

// GetBeginningOfWeek 获取一周的开始时间
func GetBeginningOfWeek(t time.Time) time.Time {
	year, month, day := t.Date()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return time.Date(year, month, day-weekday+1, 0, 0, 0, 0, t.Location())
}

// GetBeginningOfNextWeek 获取下一周的开始时间
func GetBeginningOfNextWeek(t time.Time) time.Time {
	year, month, day := t.Date()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return time.Date(year, month, day+8-weekday, 0, 0, 0, 0, t.Location())
}
