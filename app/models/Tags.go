package models

type Tags struct {
	Id        int    `json:"id" xorm:"pk autoincr"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	CreatedAt int64  `json:"created_at" xorm:"created"`
	UpdatedAt int64  `json:"updated_at" xorm:"created"`
}
