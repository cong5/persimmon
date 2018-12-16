package service

import (
	"github.com/garyburd/redigo/redis"
	"github.com/revel/revel"
)

var authService, AuthS *AuthService
var userService, UserS *UserService
var categoryService, CategoryS *CategoryService
var optionService, OptionS *OptionService
var linkService, LinkS *LinkService
var commentService, commentS *CommentService
var tagService, TagS *TagService
var postService, PostS *PostService
var postTagsService, PostTagsS *PostTagsService
var uploadsService, UploadsS *UploadsService
var baiduFanyiService, baiduFanyiS *BaiduFanyiService
var dingdingService, dingdingS *DingdingService
var redisObj redis.Conn

func InitService() {
	AuthS = &AuthService{}
	UserS = &UserService{}
	CategoryS = &CategoryService{}
	OptionS = &OptionService{}
	LinkS = &LinkService{}
	commentS = &CommentService{}
	TagS = &TagService{}
	PostS = &PostService{}
	PostTagsS = &PostTagsService{}
	UploadsS = &UploadsService{}
	baiduFanyiS = &BaiduFanyiService{}
	dingdingS = &DingdingService{}

	categoryService = CategoryS
	userService = UserS
	authService = AuthS
	optionService = OptionS
	linkService = LinkS
	commentService = commentS
	tagService = TagS
	postService = PostS
	postTagsService = PostTagsS
	uploadsService = UploadsS
	baiduFanyiService = baiduFanyiS
	dingdingService = dingdingS
}

func InitRedis() {
	openRedis := revel.Config.StringDefault("cache.redis", "false")
	redisHost := revel.Config.StringDefault("cache.hosts", "")
	print(openRedis)
	if openRedis == "false" {
		return
	}

	redis, err := redis.Dial("tcp", redisHost)
	if err != nil {
		revel.AppLog.Errorf("Dial Error %s", err)
		return
	}

	redisObj = redis
}
