package wta

import (
	"fmt"
	"time"
)

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
	anno.Hour = int8(SecondsToday / 3600)
	anno.Minute = int8(SecondsToday % 3600 / 60)
	anno.Second = int8(SecondsToday % 3600 % 60)
	return
}

// GetAnnoStr 返回世界树纪元文字表示
func (anno *Anno) GetAnnoStr() (annoStr string) {
	annoStr = anno.YearStr + anno.MonthStr + anno.DayStr
	annoStr = fmt.Sprintf(`%s　%d:%0*d:%0*d`, annoStr, anno.Hour, 2, anno.Minute, 2, anno.Second)
	annoStr = fmt.Sprintf(`%s　%s`, annoStr, anno.ChordStr)
	return
}

// GetAnnoStr 返回世界树纪元文字表示，琴弦单独返回
func (anno *Anno) GetAnnoStrSplit() (annoStr string, ChordStr string) {
	annoStr = anno.YearStr + anno.MonthStr + anno.DayStr
	annoStr = fmt.Sprintf(`%s　%d:%0*d:%0*d`, annoStr, anno.Hour, 2, anno.Minute, 2, anno.Second)
	ChordStr = anno.ChordStr
	return
}
