package models

type Comments struct {
	Id        int    `json:"id" xorm:"pk autoincr"`
	Pid       int    `json:"pid"`
	PostsId   int    `json:"posts_id"`
	Title     string `json:"title" xorm:"-"`
	Slug      string `json:"slug" xorm:"-"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Md5       string `json:"md5" xorm:"-"`
	Url       string `json:"url"`
	Content   string `json:"content"`
	Markdown  string `json:"markdown"`
	Ip        uint32 `json:"ip"`
	CreatedAt int64  `json:"created_at" xorm:"created"`
	UpdatedAt int64  `json:"updated_at" xorm:"updated"`
	Status    int    `json:"status"`
}
