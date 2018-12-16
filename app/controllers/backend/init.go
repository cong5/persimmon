package backend

import (
	"github.com/cong5/persimmon/app/service"
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/models"
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
}

func AuthInterceptor(c *revel.Controller) revel.Result {
	// 验证是否已登录
	if userId, ok := c.Session["UserId"]; ok && userId != "" {
		return nil // 已登录
	}

	// 没有登录, 判断是否是ajax操作
	if c.Request.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		return c.RenderJSON(models.Res{Status: 401, Info: "Unauthorized"})
	}

	return c.Redirect("/login")
}

func init() {
	//服务
	InitService()
	// interceptor
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &Category{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &Comment{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &File{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &Link{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &Navigation{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &Option{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &Posts{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &Settings{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &Tags{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &Trash{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &User{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &Utils{})
}
