package models

type BaiduFanyi struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	From      string `json:"from"`
	To        string `json:"to"`
	TransResult []struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	} `json:"trans_result"`
}

type MetaInfo struct {
	Posts     int `json:"posts"`
	Comments  int `json:"comments"`
	PostTrash int `json:"post_trash"`
	UserViews int `json:"user_views"`
}

type PagingContent struct {
	Data        interface{} `json:"data"`
	Total       int         `json:"total"`
	TotalPage   int         `json:"total_page"`
	CurrentPage int         `json:"current_page"`
}

type SiteMap struct {
	Slug      string `json:"slug"`
	CreatedAt int64  `json:"created_at" xorm:"created"`
}
