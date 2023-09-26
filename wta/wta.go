package wta

import (
	"strconv"
	"time"
)

const (
	wtaDay    = 85653 * time.Second
	kittenDay = `2017-04-25 00:00:00`
)

var (
	kittenTime, err = time.Parse(time.DateTime, kittenDay)
	a               Anno
	ds              uint64 // 天数戳

)

// GetAnno 返回世界树纪元
func GetAnno() (Anno, error) {
	t := time.Since(kittenTime) // 地球时间
	a.day.calendar.stamp = uint64(72 * float64(t) / float64(wtaDay))
	if a.day.calendar.stamp != ds {
		a.toAnno()
		ds = a.day.calendar.stamp
	}
	secondsToday := int32((72 * (t % wtaDay) % wtaDay).Seconds())
	a.second = uint8(secondsToday % 3600 % 60)
	a.minute = uint8(secondsToday % 3600 / 60)
	a.hour = uint8(secondsToday / 3600)
	return a, err
}

// GetYear 获取年
func (a *Anno) GetYear() year {
	return a.year
}

// GetMonth 获取月
func (a *Anno) GetMonth() month {
	return a.month
}

// GetDay 获取日
func (a *Anno) GetDay() day {
	return a.day
}

// GetStr 返回世界树纪元文字表示
func (a *Anno) GetStr() string {
	return a.year.str + a.month.str + a.day.calendar.str
}

// GetChord 返回世界树纪元琴弦
func (a *Anno) GetChord() string {
	return a.chord.str
}

// GetElementalAndImagery 获取月份的代表元灵及其意象
func (m month) GetElementalAndImagery() (string, string) {
	return m.elemental, m.imagery
}

// GetElementalAndImageryStr 获取月份的代表元灵及其意象字符串表示
func (m month) GetElementalAndImageryStr() string {
	return `～` + m.elemental + `元灵之` + m.imagery + `～`
}

// GetFlower 获取月份的代表花卉
func (m month) GetFlower() string {
	return m.flower
}

// String 实现 Stringer 接口
func (c calendar[T]) String() string {
	return c.str
}

// String 实现 Stringer 接口
func (y year) String() string {
	return y.str
}

// String 实现 Stringer 接口
func (c chord) String() string {
	return c.str
}

// String 实现 Stringer 接口
func (t *Anno) String() string {
	var (
		h = strconv.FormatUint(uint64(t.hour), 10)
		m = strconv.FormatUint(uint64(t.minute), 10)
		s = strconv.FormatUint(uint64(t.second), 10)
	)
	switch {
	case 10 > t.minute && 10 > t.second:
		return h + `:0` + m + `:0` + s
	case 10 > t.minute && 10 <= t.second:
		return h + `:0` + m + `:` + s
	case 10 <= t.minute && 10 > t.second:
		return h + `:` + m + `:0` + s
	case 10 <= t.minute && 10 <= t.second:
		return h + `:` + m + `:` + s
	default:
		return ``
	}
	// time.Date(0, 0, 0, int(t.hour), int(t.minute), int(t.second), 0, time.UTC).Format(`:04:05`)
	// fmt.Sprintf(`%d:%0*d:%0*d`, t.hour, 2, t.minute, 2, t.second)
}
