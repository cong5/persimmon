package backend

import (
	"github.com/cong5/persimmon/app/models"
	"github.com/revel/revel"
	"gopkg.in/russross/blackfriday.v2"
	"github.com/cong5/persimmon/app/utils"
)

type Posts struct {
	BaseController
}

func (c Posts) Index(categoryId int, keywords string, rows int, page int) revel.Result {
	lists := postService.GetListPaging(categoryId, keywords, rows, page, false)
	return c.RenderJSON(models.Res{Status: 200, List: lists})
}

func (c Posts) Show(id int) revel.Result {
	post, err := postService.GetPostById(id, false)
	if err != nil {
		return c.ResponseError(501, err.Error())
	}

	return c.RenderJSON(models.Res{Status: 200, Item: post})
}

func (c Posts) Store(post *models.Post) revel.Result {
	//Validation
	c.Validation.Required(post.Title)
	c.Validation.Required(post.Slug)
	c.Validation.Required(post.Markdown)
	if c.Validation.HasErrors() {
		return c.RenderJSON(models.Res{Status: 501, Info: c.Validation.Errors})
	}

	//save
	htmlContent := blackfriday.Run([]byte(post.Markdown))
	clientIP := utils.Ip2long(c.ClientIP)
	posts := models.Posts{Title: post.Title,
		Slug: post.Slug,
		Thumb: post.Thumb,
		CategoryId: post.CategoryId,
		UserId: post.UserId,
		Markdown: post.Markdown,
		Content: string(htmlContent),
		Ip: clientIP}

	postId, err := postService.Save(posts)

	if err != nil {
		return c.ResponseError(500, "add failed.")
	}

	tagService.Save(int(postId), post.Tags)
	return c.ResponseSuccess("add success.")
}

func (c Posts) Update(post models.Post) revel.Result {
	//save
	htmlContent := blackfriday.Run([]byte(post.Markdown))
	clientIP := utils.Ip2long(c.ClientIP)
	posts := models.Posts{Title: post.Title,
		Slug: post.Slug,
		Thumb: post.Thumb,
		CategoryId: post.CategoryId,
		UserId: post.UserId,
		Markdown: post.Markdown,
		Content: string(htmlContent),
		Ip: clientIP}

	_, err := postService.Update(post.Id, posts)

	if err != nil {
		return c.ResponseError(500, "Update failed.")
	}

	//Update post tags
	tagService.Save(int(post.Id), post.Tags)
	return c.ResponseSuccess("Update success.")
}

func (c Posts) Destroy(ids []int) revel.Result {
	_, err := postService.Trash(ids)
	if err != nil {
		return c.ResponseError(500, "Delete failed.")
	}

	return c.ResponseSuccess("Delete success.")
}
