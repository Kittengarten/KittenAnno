package wta

const (
	寂月 uint8 = iota
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

const (
	折纸 uint8 = 1 + iota
	赤空
	玉兰
	水光
	风荧
	玄冰
	月海
	日珥
	星灯
)

var (
	monthInfoMap = map[uint8]monthInfo{
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
	chordStrMap = map[uint8]string{
		折纸: `折纸`,
		赤空: `赤空`,
		玉兰: `玉兰`,
		水光: `水光`,
		风荧: `风荧`,
		玄冰: `玄冰`,
		月海: `月海`,
		日珥: `日珥`,
		星灯: `星灯`,
	}
)

type (
	// Year  世界树纪元的年
	year struct {
		calendar[uint64]
		IsCommon bool
	}

	// Month 世界树纪元的月
	month struct {
		elemental, imagery, flower string // 月份的代表元灵及其意象、花卉
		calendar[uint8]
		IsCommon bool
	}

	// Day 世界树纪元的月份信息
	monthInfo struct {
		str, elemental, imagery, flower string // 月份的文字表示、代表元灵及其意象、花卉
	}

	// 世界树纪元的日
	day struct {
		calendar[uint8]
		chord
	}

	// 世界树纪元琴弦
	chord struct {
		str    string // 文字
		number uint8  // 数字
	}

	calendar[T number] struct {
		str    string // 文字
		stamp  uint64 // 时间戳
		number T      // 数字
	}

	number interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}

	// Anno 世界树纪元的完整表示
	Anno struct {
		year                       // 年
		month                      // 月
		day                        // 日
		hour, minute, second uint8 // 时、分、秒的数字表示
	}

	// // 世界树纪元接口
	// Anno interface {
	// 	GetStrSplit() (annoStr, chordStr string)
	// 	GetYear() year
	// 	GetMonth() month
	// 	GetDay() day
	// }
)
