package models

type Users struct {
	Id        int    `json:"id" xorm:"pk autoincr"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at" xorm:"created"`
	UpdatedAt int64  `json:"updated_at" xorm:"updated"`
}
