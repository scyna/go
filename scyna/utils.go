package scyna

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

const microSecondPerDay = 24 * 60 * 60 * 1000000

func GetDayByTime(time time.Time) int {
	return int(time.UnixMicro() / microSecondPerDay)
}

func GetMinuteByTime(time time.Time) int64 {
	return time.Unix() / 60
}

func GetHourByTime(time time.Time) int64 {
	return time.Unix() / (60 * 60)
}

var pathrgxp = regexp.MustCompile(`:[A-z,0-9,$,-,_,.,+,!,*,',(,),\\,]{1,}`)

func PublishURL(urlPath string) string {
	ret := strings.Replace(urlPath, "/", ".", -1)
	ret = fmt.Sprintf("API%s", ret)
	return ret
}

func SubscriberURL(urlPath string) string {
	subURL := pathrgxp.ReplaceAllString(urlPath, "*")
	subURL = strings.Replace(subURL, "/", ".", -1)
	subURL = fmt.Sprintf("API%s", subURL)
	return subURL
}

func DateFromInt(timestamp uint64) string {
	return time.UnixMicro(int64(timestamp)).Format(time.RFC3339)
}
