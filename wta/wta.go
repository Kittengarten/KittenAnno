package wta

import "time"

const secondsPerDay = 85653

// GetAnno 返回世界树纪元
func GetAnno() Anno {
	var day Day = Day(time.Now().Unix() / secondsPerDay)
	return day.toAnno()
}
