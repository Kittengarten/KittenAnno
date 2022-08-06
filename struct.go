package wta

type (
	// Year 年数戳
	Year int64
	// Month 月数戳
	Month int64
	// Day 天数戳
	Day int64
	// Number64 是一个 int64
	Number64 int64
	// Number 是一个 int
	Number int
	// Luna 世界树纪元月份
	Luna int
)

// Anno 世界树纪元的时间结构体
type Anno struct {
	yearNumber  int64  // 年份的数字表示
	monthNumber int    // 月份的数字表示
	date        int    // 日期的数字表示
	yearStr     string // 年份的文字表示
	monthStr    string // 月份的文字表示
	dayStr      string // 日期的文字表示
	month       Luna   // 月份本身
}
