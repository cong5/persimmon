package models

type Options struct {
	Id        int    `json:"id" xorm:"pk autoincr"`
	Title     string `json:"title"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	Group     string `json:"group"`
	Remark    string `json:"remark"`
	Status    string `json:"status"`
	DataType  string `json:"data_type"`
	CreatedAt int64  `json:"created_at" xorm:"created"`
	UpdatedAt int64  `json:"updated_at" xorm:"updated"`
}
