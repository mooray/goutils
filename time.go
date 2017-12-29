/*
Time util
*/

package goutils

import (
	"regexp"
	"strings"
	"time"
)

//	time formatting output
//	formation supports:
//	1."2016-05-05 13:15:24"
//	2."2016年3月15日 13时15分24秒"
func TimeFormat(t *time.Time, pattern string) string {
	b, err := regexp.MatchString("^[0-9]{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}$", pattern)
	if err == nil && b {
		return t.Format("2006-01-02 15:04:05")
	}

	return ""
}

// translate month to chinese
func timeMonthsTranslate(en string) string {

	en = strings.ToLower(en)

	mon := ""
	switch en {
	case "january":
		mon = "一"
	case "february":
		mon = "二"
	case "march":
		mon = "三"
	case "april":
		mon = "四"
	case "may":
		mon = "五"
	case "june":
		mon = "六"
	case "july":
		mon = "七"
	case "august":
		mon = "八"
	case "september":
		mon = "九"
	case "october":
		mon = "十"
	case "november":
		mon = "十一"
	case "december":
		mon = "十二"
	default:
		return ""
	}

	return mon + "月"
}

// translate weekday to chinese
func _TranToCN(en string) string {
	week := ""
	switch en {
	case "monday":
		week = "一"
	case "tuesday":
		week = "二"
	case "wednesday":
		week = "三"
	case "thursday":
		week = "四"
	case "friday":
		week = "五"
	case "saturday":
		week = "六"
	case "sunday":
		week = "日"
	default:
		return ""
	}

	return "周" + week
}
