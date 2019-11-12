package handlers

import (
	"time"
)

func GetBJToday() time.Time{
	var timeLocation *time.Location
	timeLocation, _ = time.LoadLocation("Asia/Shanghai")
	var today time.Time
	today = time.Now()
	today.In(timeLocation)
	return today 
}