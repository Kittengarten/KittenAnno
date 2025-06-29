package wta

import (
	"strconv"
	"sync"
	"time"
)

const (
	ratio            = 72
	secondsPerHour   = int32(time.Hour / time.Second)
	secondsPerMinute = int32(time.Minute / time.Second)
	wtaDay           = 85653 * time.Second
	kittenDay        = `2017-04-25 00:00:00`
)

var (
	KittenTime, err = time.Parse(time.DateTime, kittenDay)
	a               Anno
	ds              uint64 // 天数戳
	mu              sync.Mutex
)

// GetAnno 返回世界树纪元
func GetAnno() (Anno, error) {
	t := time.Since(KittenTime) // 地球时间
	mu.Lock()
	defer mu.Unlock()
	a.day.stamp = uint64(ratio * float64(t) / float64(wtaDay))
	if a.day.stamp != ds {
		a.compute()
		ds = a.day.stamp
	}
	secondsToday := int32((ratio * (t % wtaDay) % wtaDay).Seconds())
	a.second = uint8(secondsToday % secondsPerHour % secondsPerMinute)
	a.minute = uint8(secondsToday % secondsPerHour / secondsPerMinute)
	a.hour = uint8(secondsToday / secondsPerHour)
	return a, err
}

// GetYear 获取年
func (a *Anno) Year() year {
	return a.year
}

// GetMonth 获取月
func (a *Anno) Month() month {
	return a.month
}

// GetDay 获取日
func (a *Anno) Day() day {
	return a.day
}

// GetDateStr 获取世界树纪元完整日期的文字表示
func (a *Anno) DateStr() string {
	return a.year.str + a.month.str + a.day.calendar.str
}

// GetChord 获取世界树纪元琴弦
func (a *Anno) Chord() string {
	return a.chord.str
}

// GetElementalAndImagery 获取月份的代表元灵及其意象
func (m month) ElementalAndImagery() (string, string) {
	return m.elemental, m.imagery
}

// GetElementalAndImageryStr 获取月份的代表元灵及其意象字符串表示
func (m month) ElementalAndImageryStr() string {
	return `～` + m.elemental + `元灵之` + m.imagery + `～`
}

// GetFlower 获取月份的代表花卉
func (m month) Flower() string {
	return m.flower
}

// String 实现 Stringer 接口，获取日期的文字表示
func (c calendar[T]) String() string {
	return c.str
}

// String 实现 Stringer 接口，获取年份的文字表示
func (y year) String() string {
	return y.str
}

// String 实现 Stringer 接口，获取琴弦的文字表示
func (c chord) String() string {
	return c.str
}

// String 实现 Stringer 接口，获取时间的文字表示
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
