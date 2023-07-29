package wta

import (
	"fmt"
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
	monthInfo                = map[Luna]MonthInfo{
		寂月: {`寂月`, `死亡`, `祈歌`, `烟花`},
		雪月: {`雪月`, `风雪`, `飘荡`, `山茶`},
		海月: {`海月`, `海洋`, `深沉`, `金花茶`},
		夜月: {`夜月`, `暗夜`, `虚乏`, `墨兰`},
		彗月: {`彗月`, `流星`, `陨落`, `腊梅`},
		凉月: {`凉月`, `寒冰`, `凝聚`, `迷迭香`},
		芷月: {`芷月`, `凛冬`, `休憩`, `茶花`},
		茸月: {`茸月`, `河流`, `苏醒`, `春兰`},
		雨月: {`雨月`, `雨露`, `降临`, `油菜花`},
		花月: {`花月`, `繁花`, `盛开`, `拟南芥`},
		梦月: {`梦月`, `梦幻`, `轨迹`, `郁金香`},
		音月: {`音月`, `韵律`, `共鸣`, `风信子`},
		晴月: {`晴月`, `云朵`, `弥散`, `紫罗兰`},
		岚月: {`岚月`, `和春`, `离去`, `鸢尾`},
		萝月: {`萝月`, `生命`, `吟唱`, `矢车菊`},
		苏月: {`苏月`, `森林`, `幽郁`, `虞美人`},
		茜月: {`茜月`, `田野`, `丰饶`, `栀子`},
		梨月: {`梨月`, `明昼`, `迷离`, `薰衣草`},
		荷月: {`荷月`, `湖泊`, `静谧`, `莲花`},
		茶月: {`茶月`, `火焰`, `灼烈`, `满天星`},
		茉月: {`茉月`, `炎夏`, `告别`, `茉莉`},
		铃月: {`铃月`, `城市`, `回响`, `紫菀`},
		信月: {`信月`, `星辰`, `守序`, `桔梗`},
		瑶月: {`瑶月`, `时间`, `归来`, `素馨`},
		风月: {`风月`, `天空`, `呓语`, `桂花`},
		叶月: {`叶月`, `大地`, `呼唤`, `芙蓉`},
		霜月: {`霜月`, `山脉`, `厚重`, `菊花`},
		奈月: {`奈月`, `清秋`, `消逝`, `油茶`},
	}
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
	for i := 1; monthCycle > i; i++ {
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
func (month Month) getYearMonth() (year Year, monthNumber int) {
	var (
		yearCycleCount = month / yearCycleMonthCount // 闰年周期数
		netMonth       = month % yearCycleMonthCount // 余下的不足一个周期的月数
		i              = 0                           // 循环次数
	)
	for yearCycle > i && netMonth >= yearCycleFirstmonthMonth[i] {
		i++
	}
	year = Year(int(yearCycleCount)*yearCycle + i - 1)              // 年数戳
	monthNumber = int(netMonth - yearCycleFirstmonthMonth[i-1] + 1) // 月份
	// 如果是闰年，月份序号整体减少 1
	if !year.isCommonYear() {
		monthNumber--
	}
	return
}

// 输出天数戳对应的月数戳、日期
func (day Day) getMonthDay() (month Month, date int) {
	var (
		monthCycleCount = day / monthCycleDayCount // 大月周期数
		netDay          = day % monthCycleDayCount // 余下的不足一个周期的天数
		i               = 0                        // 循环次数
	)
	for monthCycle > i && netDay >= monthCycleFirstdayDay[i] {
		i++
	}
	month = Month(int(monthCycleCount)*monthCycle + i - 1) // 月数戳
	date = int(netDay - monthCycleFirstdayDay[i-1] + 1)    // 日期
	return
}

// 将天数戳转换为完整的时间
func (day Day) toAnno() (anno Anno) {
	var (
		month, date             = day.getMonthDay()
		year, monthNumber       = month.getYearMonth()
		yearNumber        int64 = int64(year) + 1
	)
	anno.YearNumber = yearNumber
	anno.MonthNumber = monthNumber
	anno.Date = date
	anno.YearStr = Number64(yearNumber).getYearString()
	anno.MonthInfo = Number(monthNumber).getMonth()
	anno.DayStr = Number(date).getDate()
	return
}

// 将数字转换为中文数字
func (number Number) toString() string {
	if 0 <= number && 9 >= number {
		return numberString[3*number : 3*number+3]
	}
	return ``
}

// 将年份数字转换为年份字符串
func (number Number64) getYearString() string {
	var (
		yearLength        = len(strconv.FormatInt(int64(number), 10))
		yearConvertMemory = make([][]string, yearLength) // 第一维表示位，第二维表示内容（0 为数字原文，1 为转换后的内容）
		returnValue       = ``
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
		return `世界树纪元元年`
	}
	return fmt.Sprintf("世界树纪元%s年", returnValue)
}

// 将月份数字转换为月份信息
func (number Number) getMonth() MonthInfo {
	return monthInfo[Luna(number)]

}

// 将日期数字转换为日期字符串
func (number Number) getDate() string {
	var dayConvertMemory = make([][]string, 2) // 第一维表示位，第二维表示内容（0 为数字原文，1 为转换后的内容）
	for i := range dayConvertMemory {
		dayConvertMemory[i] = make([]string, 2)
	}
	dayConvertMemory[1][0] = strconv.Itoa(int(number) / 10)
	dayConvertMemory[0][0] = strconv.Itoa(int(number) % 10)
	switch dayConvertMemory[1][0] {
	case `0`:
		dayConvertMemory[1][1] = `初`
	case `1`:
		dayConvertMemory[1][1] = `十`
	case `2`:
		dayConvertMemory[1][1] = `廿`
	default:
		dayConvertMemory[1][1] = ``
	}
	dayConvertMemory[0][1] = (number % 10).toString()
	switch number {
	case 10:
		return `初十`
	case 20:
		return `二十`
	default:
		return dayConvertMemory[1][1] + dayConvertMemory[0][1]
	}
}
