package models

type Links struct {
	Id        int    `json:"id" xorm:"pk autoincr"`
	Name      string `json:"name"`
	Logo      string `json:"logo"`
	Group     string `json:"group"`
	Url       string `json:"url"`
	Ip        uint32 `json:"ip"`
	CreatedAt int64  `json:"created_at" xorm:"created"`
	UpdatedAt int64  `json:"updated_at" xorm:"updated"`
}
