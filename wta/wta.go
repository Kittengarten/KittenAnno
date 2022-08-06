package wta

import "time"

const (
	secondsPerDay = 85653
	kittenDay     = "2017年04月25日"
)

// GetAnno 返回世界树纪元
func GetAnno() (anno Anno, err error) {
	var (
		unix                 = time.Now().Unix()
		kittenTime, err0     = time.Parse("2006年01月02日", kittenDay)
		wtaUnix              = unix - kittenTime.Unix()
		day              Day = Day(wtaUnix / secondsPerDay) // 天数戳
		SecondsToday         = int(wtaUnix % secondsPerDay) // 当天经过的秒数
	)
	err = err0
	anno = day.toAnno()
	anno.Hour = SecondsToday / 3600
	anno.Minute = (SecondsToday % 3600) / 60
	anno.Second = (SecondsToday % 3600) % 60
	return
}
