package models

type Posts struct {
	Id         int       `json:"id" xorm:"pk autoincr"`
	Slug       string    `json:"slug"`
	Title      string    `json:"title"`
	Thumb      string    `json:"thumb"`
	CategoryId int       `json:"categoryId"`
	Categories Categorys `json:"categories" xorm:"-"`
	UserId     int       `json:"user_id"`
	Content    string    `json:"content"`
	Markdown   string    `json:"markdown"`
	Views      int       `json:"views"`
	Comments   int       `json:"comments"`
	Ip         uint32    `json:"ip"`
	Tags       []Tags    `json:"tags" xorm:"-"`
	CreatedAt  int64     `json:"created_at" xorm:"created"`
	UpdatedAt  int64     `json:"updated_at" xorm:"updated"`
	DeletedAt  int64     `json:"deleted_at"`
}

type Post struct {
	Id         int
	Slug       string
	Title      string
	Thumb      string
	CategoryId int
	UserId     int
	Content    string
	Markdown   string
	Tags       []string
}
