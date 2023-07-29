package wta

import "time"

const (
	secondsPerDay = 85653
	kittenDay     = "2017年04月25日"
)

// GetAnno 返回世界树纪元
func GetAnno() (anno Anno, err error) {
	kittenTime, err := time.Parse("2006年01月02日", kittenDay)
	var (
		unix             = time.Now().Unix()
		wtaUnix          = 72*(unix-kittenTime.Unix()) + time.Now().UnixNano()%1000000000*72/1000000000
		day          Day = Day(wtaUnix / secondsPerDay) // 天数戳
		SecondsToday     = int(wtaUnix % secondsPerDay) // 当天经过的秒数
	)
	anno = day.toAnno()
	anno.Hour = SecondsToday / 3600
	anno.Minute = (SecondsToday % 3600) / 60
	anno.Second = (SecondsToday % 3600) % 60
	return
}
