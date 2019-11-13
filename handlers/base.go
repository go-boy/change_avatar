package handlers

import (
	"time"
)

// GetBJToday 获取北京时间今天的日期
func GetBJToday() time.Time{
	var timeLocation *time.Location
	timeLocation, _ = time.LoadLocation("Asia/Shanghai")
	var today time.Time
	today = time.Now()
	today.In(timeLocation)
	return today 
}