package home

import (
	"github.com/cong5/persimmon/app/service"
	"github.com/cong5/persimmon/app/utils"
	"github.com/revel/revel"
	"html/template"
)

var authService, AuthS *service.AuthService
var userService, UserS *service.UserService
var categoryService, CategoryS *service.CategoryService
var optionService, OptionS *service.OptionService
var linkService, LinkS *service.LinkService
var commentService, commentS *service.CommentService
var tagService, TagS *service.TagService
var postService, PostS *service.PostService
var postTagsService, PostTagsS *service.PostTagsService
var uploadsService, UploadsS *service.UploadsService
var baiduFanyiService, baiduFanyiS *service.BaiduFanyiService
var notificationService, notificationS *service.NotificationService
var postTrashService, postTrashS *service.PostTrashService
var navigationService, navigationS *service.NavigationService
var feedService, FeedS *service.FeedService
var sitemapService, SitemapS *service.SitemapService
var dingdingService, dingdingS *service.DingdingService

func InitService() {
	AuthS = &service.AuthService{}
	UserS = &service.UserService{}
	CategoryS = &service.CategoryService{}
	OptionS = &service.OptionService{}
	LinkS = &service.LinkService{}
	commentS = &service.CommentService{}
	TagS = &service.TagService{}
	PostS = &service.PostService{}
	postTrashS = &service.PostTrashService{}
	PostTagsS = &service.PostTagsService{}
	UploadsS = &service.UploadsService{}
	baiduFanyiS = &service.BaiduFanyiService{}
	notificationS = &service.NotificationService{}
	navigationS = &service.NavigationService{}
	FeedS = &service.FeedService{}
	SitemapS = &service.SitemapService{}
	dingdingS = &service.DingdingService{}

	categoryService = CategoryS
	userService = UserS
	authService = AuthS
	optionService = OptionS
	linkService = LinkS
	commentService = commentS
	tagService = TagS
	postService = PostS
	postTrashService = postTrashS
	postTagsService = PostTagsS
	uploadsService = UploadsS
	baiduFanyiService = baiduFanyiS
	notificationService = notificationS
	navigationService = navigationS
	feedService = FeedS
	sitemapService = SitemapS
	dingdingService = dingdingS
}

func NewTemplateFunc() {

	revel.TemplateFuncs["date"] = func(format string, timestamp int64) template.JS {
		if timestamp == 0 {
			return ""
		}
		date := utils.Date(format, timestamp)
		return template.JS(date)
	}
	revel.TemplateFuncs["isEmpty"] = func(arg string) bool {
		if arg == "" {
			return true
		}
		return false
	}

}

func init() {
	//服务
	InitService()
	NewTemplateFunc()
}
