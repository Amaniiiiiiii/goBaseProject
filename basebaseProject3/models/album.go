package models

//结构体名称和结构体中的字段都是大写
//下划线省略 AddTime 对应 add_time

type Album struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

//默认表名是结构体的复数形式Albums 也可以自定义
// 我们自定义Album 对应 album 表
func (Album) TableName() string {
	return "album"
}
