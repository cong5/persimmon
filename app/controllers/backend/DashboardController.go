package backend

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/models"
)

type Dashboard struct {
	BaseController
}

func (c Dashboard) Index() revel.Result {
	siteUrl := c.Request.Host
	siteName := "Persimmon"
	return c.Render(siteUrl, siteName)
}

func (c Dashboard) MetaCount() revel.Result {
	post, pErr := postService.CountPost(0,"")
	if pErr != nil {
		post = 0
	}

	postTrash, ptErr := postTrashService.CountTrashPost()
	if ptErr != nil {
		postTrash = 0
	}

	comment, cErr := commentService.CountComment(0)
	if cErr != nil {
		comment = 0
	}

	viewsCount, vErr := postService.SumViews()
	if vErr != nil {
		viewsCount = 0
	}

	metaInfo := models.MetaInfo{Posts: post,
		PostTrash: postTrash,
		Comments: comment,
		UserViews: viewsCount}
	return c.RenderJSON(models.Res{Status: 200, Item: metaInfo})
}
