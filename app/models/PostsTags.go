package models

type PostsTags struct {
	PostsId   int   `json:"posts_id"`
	TagsId    int   `json:"tags_id"`
	CreatedAt int64 `json:"created_at" xorm:"created"`
	UpdatedAt int64 `json:"updated_at" xorm:"updated"`
}
