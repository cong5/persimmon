package models

type Categorys struct {
	Id          int    `json:"id" xorm:"pk autoincr"`
	Name        string `json:"name"`
	Pid         int    `json:"pid"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Ip          uint32 `json:"ip"`
	CreatedAt   int64  `json:"created_at" xorm:"created"`
	UpdatedAt   int64  `json:"updated_at" xorm:"updated"`
}
