package wta

type (
	// Year 年数戳
	Year int64
	// Month 月数戳
	Month int64
	// Day 天数戳
	Day int64
	// Annual 世界树纪元年份
	Annual int64
	// Luna 世界树纪元月份
	Luna int8
	// Date 世界树纪元日期
	Date int8
	// Chord 世界树纪元琴弦
	Chord int8
	// Number 是一个可转换为中文的数字
	Number[T Annual | Date] int64
)

// Anno 世界树纪元的时间结构体
type Anno struct {
	YearNumber  int64  // 年份的数字表示
	MonthNumber int8   // 月份的数字表示
	Date        int8   // 日期的数字表示
	YearStr     string // 年份的文字表示
	MonthInfo          // 月份信息结构体
	DayStr      string // 日期的文字表示
	ChordStr    string // 琴弦的文字表示
	Hour        int8   // 时
	Minute      int8   // 分
	Second      int8   // 秒
}

// MonthInfo 世界树纪元的月份信息结构体
type MonthInfo struct {
	MonthStr  string // 月份的文字表示
	Elemental string // 月份的代表元灵
	Imagery   string // 月份的代表元灵之意象
	Flower    string // 月份的代表花卉
}
