package wta

import (
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
	yearCycleFirstmonthMonth [yearCycle]uint16  // 闰年周期中，每年的首月所处的月数戳
	monthCycleFirstdayDay    [monthCycle]uint16 // 大月周期中，每月的首日所处的天数戳
)

// 计算出闰年和大月
func init() {
	yearCycleFirstmonthMonthCompute()
	monthCycleFirstdayDayCompute()
}

// 计算闰年
func yearCycleFirstmonthMonthCompute() {
	yearCycleFirstmonthMonth[0] = 0
	for i := range [yearCycle - 1]struct{}{} {
		yearCycleFirstmonthMonth[i+1] = yearCycleFirstmonthMonth[i] + commonYearMonthCount
		anno := new(Anno)
		if anno.year.stamp = uint64(i); !anno.isCommonYear() {
			// 如果是闰年，额外增加一个月
			yearCycleFirstmonthMonth[i+1]++
		}
	}
}

// 判断是否平年
func (anno *Anno) isCommonYear() bool {
	anno.year.IsCommon = !slices.Contains([]int8{1, 4, 7, 10, 13, 15, 18, 21, 24, 27}, int8(anno.year.stamp%yearCycle))
	return anno.year.IsCommon
}

// 计算大月
func monthCycleFirstdayDayCompute() {
	monthCycleFirstdayDay[0] = 0
	for i := range [monthCycle - 1]struct{}{} {
		monthCycleFirstdayDay[i+1] = monthCycleFirstdayDay[i] + commonMonthDayCount
		anno := new(Anno)
		if anno.month.stamp = uint64(i); !anno.isCommonMonth() {
			// 如果是大月，额外增加一天
			monthCycleFirstdayDay[i+1]++
		}
	}
}

// 判断是否小月
func (anno *Anno) isCommonMonth() bool {
	anno.month.IsCommon = !slices.Contains([]int8{1, 4, 8}, int8(anno.month.stamp%monthCycle))
	return anno.month.IsCommon
}

// 输出月数戳对应的年数戳，获取月对应的数字
func (anno *Anno) getYearMonth() {
	var (
		yearCycleCount = anno.month.stamp / yearCycleMonthCount         // 闰年周期数
		netMonth       = uint16(anno.month.stamp % yearCycleMonthCount) // 余下的不足一个周期的月数
		i              int                                              // 循环次数
	)
	for i = range yearCycleFirstmonthMonth {
		if netMonth < yearCycleFirstmonthMonth[i] {
			i-- // 去除多余的一次循环
			break
		}
	}
	anno.year.stamp = yearCycleCount*yearCycle + uint64(i)            // 年数戳
	anno.month.number = uint8(netMonth - yearCycleFirstmonthMonth[i]) // 月份
	// 如果是平年，月份序号整体增加 1
	if anno.isCommonYear() {
		anno.month.number++
	}
}

// 输出天数戳对应的月数戳，获取日对应的数字
func (anno *Anno) getMonthDay() {
	var (
		monthCycleCount = anno.day.stamp / monthCycleDayCount         // 大月周期数
		netDay          = uint16(anno.day.stamp % monthCycleDayCount) // 余下的不足一个周期的天数
		i               int                                           // 循环次数
	)
	for i = range monthCycleFirstdayDay {
		if netDay < monthCycleFirstdayDay[i] {
			i-- // 去除多余的一次循环
			break
		}
	}
	anno.month.stamp = monthCycleCount*monthCycle + uint64(i)             // 月数戳
	anno.day.calendar.number = 1 + uint8(netDay-monthCycleFirstdayDay[i]) // 日期
	return
}

// 将天数戳转换为完整的时间
func (anno *Anno) toAnno() {
	s := anno.month.stamp
	anno.getMonthDay()
	anno.getDate()
	anno.chord.number = 1 + uint8(anno.day.stamp%9)
	anno.chord.str = chordStrMap[anno.chord.number]
	if anno.month.stamp == s {
		return
	}
	s = anno.year.stamp
	anno.getYearMonth()
	anno.getMonth()
	if anno.year.stamp == s {
		return
	}
	anno.year.number = 1 + anno.year.stamp
	anno.getYear()
	return
}

// 将数字转换为中文数字
func parseRune[T number](n T) rune {
	if 0 <= n && 9 >= n {
		return []rune(numberString)[n]
	}
	return 0
}

// 将年份转换为年份字符串
func (anno *Anno) getYear() {
	var (
		c             int // c 为 0 表示个位，为 1 表示十位，以此类推
		l             = len(strconv.FormatInt(int64(anno.year.number), 10))
		yearConverted = make([]rune, l, l)
	)
	for i := anno.year.number; 0 < i; i /= 10 {
		v := i % 10
		yearConverted[c] = parseRune(v)
		c++
	}
	slices.Reverse(yearConverted)
	if 1 == anno.year.number {
		anno.year.str = `世界树纪元元年`
		return
	}
	anno.year.str = `世界树纪元` + string(yearConverted) + `年`
}

// 将月份转换为月份信息
func (anno *Anno) getMonth() {
	anno.month.str = monthInfoMap[anno.month.number].str
	anno.month.elemental = monthInfoMap[anno.month.number].elemental
	anno.month.imagery = monthInfoMap[anno.month.number].imagery
	anno.month.flower = monthInfoMap[anno.month.number].flower
}

// 将日期转换为日期字符串
func (anno *Anno) getDate() {
	var ten rune
	switch strconv.FormatInt(int64(anno.day.calendar.number)/10, 10) {
	case `0`:
		ten = '初'
	case `1`:
		ten = '十'
	case `2`:
		ten = '廿'
	}
	switch anno.day.calendar.number {
	case 10:
		anno.day.calendar.str = `初十`
	case 20:
		anno.day.calendar.str = `二十`
	default:
		anno.day.calendar.str = string([]rune{ten, parseRune(anno.day.calendar.number % 10)})
	}
}
