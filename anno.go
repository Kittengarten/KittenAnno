package wta

import (
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
)

const (
	寂月 Luna = iota
	雪月
	海月
	夜月
	彗月
	凉月
	芷月
	茸月
	雨月
	花月
	梦月
	音月
	晴月
	岚月
	萝月
	苏月
	茜月
	梨月
	荷月
	茶月
	茉月
	铃月
	信月
	瑶月
	风月
	叶月
	霜月
	奈月
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
	for i := 1; i <= yearCycle; i++ {
		if Year(i - 1).isCommonYear() {
			yearCycleFirstmonthMonth[i] = (yearCycleFirstmonthMonth[i-1] + commonYearMonthCount)
		} else {
			yearCycleFirstmonthMonth[i] = (yearCycleFirstmonthMonth[i-1] + commonYearMonthCount + 1)
		}
	}
}

// 判断是否平年
func (year Year) isCommonYear() bool {
	var (
		netYear = year % yearCycle
		years   = []Year{1, 4, 7, 10, 13, 15, 18, 21, 24, 27}
	)
	for _, v := range years {
		if netYear == v {
			return false
		}
	}
	return true
}

// 计算大月
func monthCycleFirstdayDayCompute() {
	monthCycleFirstdayDay[0] = 0
	for i := 1; i < monthCycle; i++ {
		if Month(i - 1).isCommonMonth() {
			monthCycleFirstdayDay[i] = (monthCycleFirstdayDay[i-1] + commonMonthDayCount)
		} else {
			monthCycleFirstdayDay[i] = (monthCycleFirstdayDay[i-1] + commonMonthDayCount + 1)
		}
	}
}

// 判断是否小月
func (month Month) isCommonMonth() bool {
	var (
		netMonth = month % monthCycle
		months   = []Month{1, 4, 8}
	)
	for _, v := range months {
		if netMonth == v {
			return false
		}
	}
	return true
}

// 输出月数戳对应的年数戳、月份
func (month Month) toYearMonth() (year Year, monthNumber int) {
	var (
		yearCycleCount = month / yearCycleMonthCount // 闰年周期数
		netMonth       = month % yearCycleMonthCount // 余下的不足一个周期的月数
		i              = 0                           // 循环次数
	)
	for netMonth >= yearCycleFirstmonthMonth[i] && i < yearCycle {
		i++
	}
	year = Year(int(yearCycleCount)*yearCycle + i - 1)            // 年数戳
	monthNumber = int(netMonth - yearCycleFirstmonthMonth[i] + 1) // 月份
	// 如果是闰年，月份序号整体减少 1
	if !year.isCommonYear() {
		monthNumber--
	}
	return
}

// 输出天数戳对应的月数戳、日期
func (day Day) toMonthDay() (month Month, date int) {
	var (
		monthCycleCount = day / monthCycleDayCount // 大月周期数
		netDay          = day % monthCycleDayCount // 余下的不足一个周期的天数
		i               = 0                        // 循环次数
	)
	for netDay >= monthCycleFirstdayDay[i] && i < monthCycle {
		i++
	}
	month = Month(int(monthCycleCount)*monthCycle + i - 1) // 月数戳
	date = int(netDay - monthCycleFirstdayDay[i] + 1)      // 日期
	return
}

// 将天数戳转换为完整的时间
func (day Day) toString() (anno Anno) {
	var (
		month, date             = day.toMonthDay()
		year, monthNumber       = month.toYearMonth()
		yearNumber        int64 = int64(year) + 1
	)
	anno.yearNumber = yearNumber
	anno.monthNumber = monthNumber
	anno.date = date
	anno.yearStr = Number64(yearNumber).toYearString()
	anno.monthStr = Number(monthNumber).toMonthString()
	anno.dayStr = Number(date).toDate()
	return
}

// 将数字转换为中文数字
func (number Number) toString() string {
	switch number {
	case 0:
		return "〇"
	case 1:
		return "一"
	case 2:
		return "二"
	case 3:
		return "三"
	case 4:
		return "四"
	case 5:
		return "五"
	case 6:
		return "六"
	case 7:
		return "七"
	case 8:
		return "八"
	case 9:
		return "九"
	default:
		return ""
	}
}

// 将年份数字转换为年份字符串
func (number Number64) toYearString() string {
	var (
		yearLength        = len(strconv.FormatInt(int64(number), 10))
		yearConvertMemory = make([][]string, yearLength) // 第一维表示位，第二维表示内容（0 为数字原文，1 为转换后的内容）
		returnValue       = ""
	)
	for i := range yearConvertMemory {
		yearConvertMemory[i] = make([]string, 2)
	}
	for i := int64(number); i > 0; i /= 10 {
		Circulate := 0 // 0 表示个位，1 表示十位，以此类推
		v := i % 10
		yearConvertMemory[Circulate][0] = strconv.FormatInt(v, 10)
		yearConvertMemory[Circulate][1] = Number(v).toString()
		returnValue = yearConvertMemory[Circulate][1] + returnValue
	}
	if number == 1 {
		return "世界树纪元元"
	}
	return "世界树纪元" + returnValue
}

// 将月份数字转换为月份字符串
func (number Number) toMonthString() string {
	switch Luna(number) {
	case 寂月:
		return "寂月"
	case 雪月:
		return "雪月"
	case 海月:
		return "海月"
	case 夜月:
		return "夜月"
	case 彗月:
		return "彗月"
	case 凉月:
		return "凉月"
	case 芷月:
		return "芷月"
	case 茸月:
		return "茸月"
	case 雨月:
		return "雨月"
	case 花月:
		return "花月"
	case 梦月:
		return "梦月"
	case 音月:
		return "音月"
	case 晴月:
		return "晴月"
	case 岚月:
		return "岚月"
	case 萝月:
		return "萝月"
	case 苏月:
		return "苏月"
	case 茜月:
		return "茜月"
	case 梨月:
		return "梨月"
	case 荷月:
		return "荷月"
	case 茶月:
		return "茶月"
	case 茉月:
		return "茉月"
	case 铃月:
		return "铃月"
	case 信月:
		return "信月"
	case 瑶月:
		return "瑶月"
	case 风月:
		return "风月"
	case 叶月:
		return "叶月"
	case 霜月:
		return "霜月"
	case 奈月:
		return "奈月"
	default:
		return ""
	}
}

// 将日期数字转换为日期字符串
func (number Number) toDate() string {
	var dayConvertMemory = make([][]string, 2) // 第一维表示位，第二维表示内容（0 为数字原文，1 为转换后的内容）
	for i := range dayConvertMemory {
		dayConvertMemory[i] = make([]string, 2)
	}
	dayConvertMemory[1][0] = strconv.Itoa(int(number) / 10)
	dayConvertMemory[0][0] = strconv.Itoa(int(number) % 10)
	switch dayConvertMemory[1][0] {
	case "0":
		dayConvertMemory[1][1] = "初"
		break
	case "1":
		dayConvertMemory[1][1] = "十"
		break
	case "2":
		dayConvertMemory[1][1] = "廿"
		break
	default:
		dayConvertMemory[1][1] = ""
	}
	dayConvertMemory[0][1] = (number % 10).toString()
	switch number {
	case 10:
		return "初十"
	case 20:
		return "二十"
	default:
		return dayConvertMemory[1][1] + dayConvertMemory[0][1]
	}
}
