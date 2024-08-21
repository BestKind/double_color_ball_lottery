package models

type Lottery struct {
	ID       int64  `gorm:"column:id" json:"id"`
	Version  string `gorm:"column:version" json:"version"`
	OpenTime string `gorm:"column:open_time" json:"open_time"`
	Week     string `gorm:"column:week" json:"week"`
	Red1     string `gorm:"column:red_1" json:"red_1"`
	Red2     string `gorm:"column:red_2" json:"red_2"`
	Red3     string `gorm:"column:red_3" json:"red_3"`
	Red4     string `gorm:"column:red_4" json:"red_4"`
	Red5     string `gorm:"column:red_5" json:"red_5"`
	Red6     string `gorm:"column:red_6" json:"red_6"`
	Blue     string `gorm:"column:blue" json:"blue"`
	L1Val    string `gorm:"column:l1_val" json:"l1_val"`
	L1Cnt    string `gorm:"column:l1_cnt" json:"l1_cnt"`
	L2Val    string `gorm:"column:l2_val" json:"l2_val"`
	L2Cnt    string `gorm:"column:l2_cnt" json:"l2_cnt"`
	L3Cnt    string `gorm:"column:l3_cnt" json:"l3_cnt"`
	Total    string `gorm:"column:total" json:"total"`
	Sale     string `gorm:"column:sale" json:"sale"`
}

func (l *Lottery) TableName() string {
	return "lottery"
}

type Grade struct {
	Type      int    `json:"type"`
	Typenum   string `json:"typenum"`
	Typemoney string `json:"typemoney"`
}

type Result struct {
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	DetailsLink string  `json:"detailsLink"`
	VideoLink   string  `json:"videoLink"`
	Date        string  `json:"date"`
	Week        string  `json:"week"`
	Red         string  `json:"red"`
	Blue        string  `json:"blue"`
	Sales       string  `json:"sales"`
	Poolmoney   string  `json:"poolmoney"`
	Content     string  `json:"content"`
	Prizegrades []Grade `json:"prizegrades"`
}

type LotteryRes struct {
	State    int      `json:"state"`
	Total    int      `json:"total"`
	Message  string   `json:"message"`
	PageNum  int      `json:"pageNum"`
	PageNo   int      `json:"pageNo"`
	PageSize int      `json:"pageSize"`
	Tflag    int      `json:"Tflag"`
	Result   []Result `josn:"result"`
}
