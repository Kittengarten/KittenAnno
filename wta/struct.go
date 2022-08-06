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
	YearNumber  int64  // 年份的数字表示
	MonthNumber int    // 月份的数字表示
	Date        int    // 日期的数字表示
	YearStr     string // 年份的文字表示
	MonthStr    string // 月份的文字表示
	DayStr      string // 日期的文字表示
	Month       Luna   // 月份本身
	Hour        int    // 时
	Minute      int    // 分
	Second      int    // 秒
}
