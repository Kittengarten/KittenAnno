package wta

import (
	"fmt"
	"slices"
	"strconv"
)

const (
	commonYearMonthCount   = 27                                                      // 平年的月数
	commonMonthDayCount    = 20                                                      // 小月的天数
	yearCycle              = 29                                                      // 闰年周期的年数
	cycleLeapYearCount     = 10                                                      // 每周期的闰年数
	monthCycle             = 10                                                      // 大月周期的月数
	cycleGreaterMonthCount = 3                                                       // 每周期的大月数
	yearCycleMonthCount    = yearCycle*commonYearMonthCount + cycleLeapYearCount     // 闰年周期的月数
	monthCycleDayCount     = monthCycle*commonMonthDayCount + cycleGreaterMonthCount // 大月周期的天数
	numberString           = `〇一二三四五六七八九`                                            // 数字对应的字符
)

var (
	yearCycleFirstmonthMonth [yearCycle]Month // 闰年周期中，每年的首月所处的月数戳
	monthCycleFirstdayDay    [monthCycle]Day  // 大月周期中，每月的首日所处的天数戳
)

// 计算出闰年和大月
func init() {
	yearCycleFirstmonthMonthCompute()
	monthCycleFirstdayDayCompute()
}

// 计算闰年
func yearCycleFirstmonthMonthCompute() {
	yearCycleFirstmonthMonth[0] = 0
	for i := 1; yearCycle > i; i++ {
		yearCycleFirstmonthMonth[i] = yearCycleFirstmonthMonth[i-1] + commonYearMonthCount
		if !Year(i - 1).isCommon() {
			// 如果是闰年，额外增加一个月
			yearCycleFirstmonthMonth[i]++
		}
	}
}

// 判断是否平年
func (year Year) isCommon() bool {
	var (
		years   = []int8{1, 4, 7, 10, 13, 15, 18, 21, 24, 27}
		netYear = int8(year % yearCycle)
	)
	return !slices.Contains(years, netYear)
}

// 计算大月
func monthCycleFirstdayDayCompute() {
	monthCycleFirstdayDay[0] = 0
	for i := 1; monthCycle > i; i++ {
		monthCycleFirstdayDay[i] = monthCycleFirstdayDay[i-1] + commonMonthDayCount
		if !Month(i - 1).isCommon() {
			// 如果是大月，额外增加一天
			monthCycleFirstdayDay[i]++
		}
	}
}

// 判断是否小月
func (month Month) isCommon() bool {
	var (
		months   = []int8{1, 4, 8}
		netMonth = int8(month % monthCycle)
	)
	return !slices.Contains(months, netMonth)
}

// 输出月数戳对应的年数戳、月份
func (month Month) getYearMonth() (year Year, monthNumber Luna) {
	var (
		yearCycleCount = month / yearCycleMonthCount // 闰年周期数
		netMonth       = month % yearCycleMonthCount // 余下的不足一个周期的月数
		i              int                           // 循环次数
	)
	for i = range yearCycleFirstmonthMonth {
		if netMonth < yearCycleFirstmonthMonth[i] {
			i-- // 去除多余的一次循环
			break
		}
	}
	year = Year(int(yearCycleCount)*yearCycle + i)                 // 年数戳
	monthNumber = Luna(netMonth - yearCycleFirstmonthMonth[i] + 1) // 月份
	// 如果是闰年，月份序号整体减少 1
	if !year.isCommon() {
		monthNumber--
	}
	return
}

// 输出天数戳对应的月数戳、日期
func (day Day) getMonthDay() (month Month, date Date) {
	var (
		monthCycleCount = day / monthCycleDayCount // 大月周期数
		netDay          = day % monthCycleDayCount // 余下的不足一个周期的天数
		i               int                        // 循环次数
	)
	for i = range monthCycleFirstdayDay {
		if netDay < monthCycleFirstdayDay[i] {
			i-- // 去除多余的一次循环
			break
		}
	}
	month = Month(int(monthCycleCount)*monthCycle + i) // 月数戳
	date = Date(netDay - monthCycleFirstdayDay[i] + 1) // 日期
	return
}

// 将天数戳转换为完整的时间
func (day Day) toAnno() (anno Anno) {
	var (
		month, date       = day.getMonthDay()
		year, monthNumber = month.getYearMonth()
		yearNumber        = Annual(year) + 1
		chordNumber       = Chord(day%9 + 1)
	)
	anno = Anno{
		AnnoStr: AnnoStr{
			YearStr:   yearNumber.getYear(),
			MonthInfo: monthNumber.getMonth(),
			DayStr:    date.getDate(),
			ChordStr:  chord[chordNumber],
		},
		YearNumber:  uint64(yearNumber),
		MonthNumber: uint8(monthNumber),
		Date:        uint8(date),
	}
	return
}

// 将数字转换为中文数字
func (n Number[T]) toRune() rune {
	if 0 <= n && 9 >= n {
		return []rune(numberString)[n]
	}
	return 0
}

// 将年份转换为年份字符串
func (a Annual) getYear() (s string) {
	var (
		c             int // c 为 0 表示个位，为 1 表示十位，以此类推
		l             = len(strconv.FormatInt(int64(a), 10))
		yearConverted = make([]rune, l, l)
	)
	for i := a; i > 0; i /= 10 {
		v := i % 10
		yearConverted[c] = Number[Annual](v).toRune()
		c++
	}
	slices.Reverse(yearConverted)
	if a == 1 {
		return `世界树纪元元年`
	}
	return fmt.Sprintf(`世界树纪元%s年`, string(yearConverted))
}

// 将月份转换为月份信息
func (m Luna) getMonth() MonthInfo {
	return monthInfo[m]

}

// 将日期转换为日期字符串
func (d Date) getDate() string {
	var dayConverted struct {
		one, ten rune
	}
	switch strconv.FormatInt(int64(d)/10, 10) {
	case `0`:
		dayConverted.ten = '初'
	case `1`:
		dayConverted.ten = '十'
	case `2`:
		dayConverted.ten = '廿'
	}
	switch dayConverted.one = Number[Date](d % 10).toRune(); d {
	case 10:
		return `初十`
	case 20:
		return `二十`
	default:
		return string([]rune{dayConverted.ten, dayConverted.one})
	}
}
