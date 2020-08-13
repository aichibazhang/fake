package util

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// 生成一段数字之间随机整数
func RandIntRangeBetween(min, max int64) int64 {
	if min >= max {
		return min
	}
	return rand.Int63n((max+1)-min) + min
}

// 最多生成18位,生成随机整数
func RandIntRangeRand(digits int64) int64 {
	return RandIntRangeBetween(int64(math.Pow(float64(10), float64(digits-1))),
		int64(math.Pow(float64(10), float64(digits)))-1)
}

//生成n位随机小数
func RandFloatRangeRand(digits int64) float64 {
	number, _ := strconv.ParseFloat(fmt.Sprintf("%.[1]*f", digits, rand.Float64()*math.Pow(float64(10), float64(digits-1))), 64)
	_, point := math.Modf(number)
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.[1]*f", digits, point), 64)
	return value
}

// 一个时间段之间的日期
func RandDate(start string, end string, flag bool) string {
	startTime, _ := time.ParseInLocation("2006-01-02", start, time.Local)
	endTime, _ := time.ParseInLocation("2006-01-02", end, time.Local)
	if startTime.After(endTime) {
		return start
	}
	date := randDate(startTime, endTime)
	var rtnDate string
	switch flag {
	case true:
		rtnDate = time.Unix(date, 0).Format("2006-01-02 15:04:05")
	case false:
		rtnDate = time.Unix(date, 0).Format("2006-01-02")
	}
	return rtnDate
}
func randDate(startTime time.Time, endTime time.Time) int64 {
	rtnDate := startTime.Unix() + RandIntRangeBetween(0, endTime.Unix()-startTime.Unix())
	if rtnDate == startTime.Unix() || rtnDate == endTime.Unix() {
		randDate(startTime, endTime)
	}
	return rtnDate
}
