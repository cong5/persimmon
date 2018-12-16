package models

type Attachments struct {
	Id        int    `json:"id" xorm:"pk autoincr"`
	Path      string `json:"path"`
	UserId    int    `json:"user_id"`
	Hash1     string `json:"hash1"`
	Md5       string `json:"md5"`
	Ip        string `json:"ip"`
	CreatedAt int64  `json:"created_at" xorm:"created"`
	UpdatedAt int64  `json:"updated_at" xorm:"updated"`
}
