package wta

import (
	"fmt"
	"time"
)

const (
	wtaDay    = 85653 * time.Second
	kittenDay = `2017-04-25 00:00:00`
)

var kittenTime, err = time.Parse(time.DateTime, kittenDay)

// GetAnno 返回世界树纪元
func GetAnno() (Anno, error) {
	var (
		t            = time.Since(kittenTime)                        // 地球时间
		day          = Day(72 * float64(t) / float64(wtaDay))        // 天数戳
		SecondsToday = int32((72 * (t % wtaDay) % wtaDay).Seconds()) // 当天经过的秒数
	)
	anno := day.toAnno()
	anno.Hour = uint8(SecondsToday / 3600)
	anno.Minute = uint8(SecondsToday % 3600 / 60)
	anno.Second = uint8(SecondsToday % 3600 % 60)
	return anno, err
}

// GetAnnoStr 返回世界树纪元文字表示
func (anno *Anno) GetAnnoStr() (annoStr string) {
	annoStr = anno.YearStr + anno.MonthStr + anno.DayStr
	annoStr = fmt.Sprintf(`%s　%d:%0*d:%0*d`, annoStr, anno.Hour, 2, anno.Minute, 2, anno.Second)
	annoStr = fmt.Sprintf(`%s　%s`, annoStr, anno.ChordStr)
	return
}

// GetAnnoStr 返回世界树纪元文字表示，琴弦单独返回
func (anno *Anno) GetAnnoStrSplit() (annoStr string, chordStr string) {
	annoStr = anno.YearStr + anno.MonthStr + anno.DayStr
	annoStr = fmt.Sprintf(`%s　%d:%0*d:%0*d`, annoStr, anno.Hour, 2, anno.Minute, 2, anno.Second)
	chordStr = anno.ChordStr
	return
}
