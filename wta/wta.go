package wta

import "time"

const secondsPerDay = 85653

// GetAnno 返回世界树纪元
func GetAnno() (anno Anno) {
	var (
		day          Day = Day(time.Now().Unix() / secondsPerDay) // 天数戳
		SecondsToday     = int(time.Now().Unix() % secondsPerDay) // 当天经过的秒数
	)
	anno = day.toAnno()
	anno.Hour = SecondsToday / 3600
	anno.Minute = (SecondsToday % 3600) / 60
	anno.Second = (SecondsToday % 3600) % 60
	return anno
}
