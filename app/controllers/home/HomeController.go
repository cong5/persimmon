package home

import (
	"github.com/cong5/persimmon/app/utils"
	"github.com/revel/revel"
)

type Home struct {
	AppController
}

func (c *Home) Index(page int) revel.Result {

	page = utils.IntDefault(page > 0, page, FirstPage)
	postArr, err := postService.SearchList(0, "", PageSize, page, false)
	if err != nil {
		c.Flash.Error(err.Error())
	}

	total, cErr := postService.CountPost(0, "")
	if cErr != nil {
		c.Flash.Error(cErr.Error())
	}

	c.ViewAssign(postArr, total, PageSize, page)
	return c.Render()
}

func (c *Home) Post(slug string) revel.Result {
	c.GetGlobalInfo()
	post, err := postService.GetPostBySlug(slug, false)
	if slug == "" || err != nil || post.Id <= 0 {
		return c.NotFound("很抱歉，没有找到这个页面.")
	}

	post.Views = postService.GetViews(post, true)
	return c.Render(post)
}

func (c *Home) Tag(name string, page int) revel.Result {
	page = utils.IntDefault(page > 0, page, 1)

	if name == "" {
		c.Flash.Error("标签不能为空.")
	}

	postIdArr, pErr := postTagsService.GetPostIdsByTagName(name)
	if pErr != nil {
		c.Flash.Error("很抱歉，根据这个标签没有找到文章.")
	}

	count := len(postIdArr)
	start := (page - 1) * PageSize
	newPostIdArr := utils.SliceIntArr(postIdArr, start, PageSize, count)

	if len(newPostIdArr) <= 0 {
		c.Flash.Error("很抱歉，找不到你想要的页面.")
	}

	postArr, postErr := postService.GetPostByIdArr(newPostIdArr, false)
	if postErr != nil {
		c.Flash.Error("很抱歉，根据这个标签没有找到文章.")
	}

	c.ViewAssign(postArr, count, PageSize, page)
	return c.Render()
}

func (c *Home) Categories(slug string, page int) revel.Result {
	category, cErr := categoryService.GetCategoryBySlug(slug, false)
	if cErr != nil {
		c.Flash.Error(cErr.Error())
		return c.NotFound("很抱歉，页面没找到")
	}

	page = utils.IntDefault(page > 0, page, 1)
	postArr, err := postService.SearchList(category.Id, "", PageSize, page, false)
	if err != nil {
		c.Flash.Error(err.Error())
	}

	total, cErr := postService.CountPost(category.Id, "")
	if cErr != nil {
		c.Flash.Error(cErr.Error())
	}

	c.ViewAssign(postArr, total, PageSize, page)
	return c.Render()
}

func (c *Home) Friends() revel.Result {
	c.GetGlobalInfo()
	links, err := linkService.GetList(9999, 1, false)
	if err != nil {
		c.Flash.Error(err.Error())
	}

	return c.Render(links)
}

func (c *Home) Feed() revel.Result {
	domain := c.Request.Host
	xml, err := feedService.BuildFeed(domain)
	if err != nil {
		c.Flash.Error(err.Error())
	}

	return myRenderXML(xml)
}

func (c *Home) SiteMap(platform string) revel.Result {
	domain := c.Request.Host
	xml, err := sitemapService.BuildSiteMap(domain, platform)
	if err != nil {
		c.Flash.Error(err.Error())
	}

	return myRenderXML(xml)
}

func (c *Home) Debug() revel.Result {
	_, err := dingdingService.SendNotice(" 测试消息")

	revel.AppLog.Debugf("%s", err)

	return c.RenderJSON("")
}
