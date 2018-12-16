package backend

import (
	"github.com/cong5/persimmon/app/models"
	"github.com/cong5/persimmon/app/utils"
	"github.com/revel/revel"
	"gopkg.in/russross/blackfriday.v2"
)

type Comment struct {
	BaseController
}

func (c Comment) Index(rows int, page int) revel.Result {
	lists, err := commentService.GetListPaging(rows, page, false)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(models.Res{Status: 200, List: lists})
}

func (c Comment) Show(id int) revel.Result {
	comment, err := commentService.GetCommentById(id, false)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(models.Res{Status: 200, Item: comment})
}

func (c Comment) Store(comment *models.Comments) revel.Result {
	//Validation
	c.Validation.Required(comment.Name)
	c.Validation.Required(comment.Url)
	if c.Validation.HasErrors() {
		return c.RenderJSON(models.Res{Status: 501, Info: c.Validation.Errors})
	}

	//save
	htmlContent := blackfriday.Run([]byte(comment.Markdown))
	clientIP := utils.Ip2long(c.ClientIP)
	comments := models.Comments{PostsId: comment.PostsId,
		Name:     comment.Name,
		Email:    comment.Email,
		Url:      comment.Url,
		Content:  string(htmlContent),
		Markdown: comment.Markdown,
		Ip:       clientIP}

	_, err := commentService.Save(comments)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("add success.")
}

func (c Comment) Update(content *models.Comments) revel.Result {

	if content.Id <= 0 {
		return c.ResponseError(501, "Params failed.")
	}

	//save
	clientIP := utils.Ip2long(c.ClientIP)
	comment := models.Comments{PostsId: content.PostsId,
		Name:     content.Name,
		Email:    content.Email,
		Url:      content.Url,
		Content:  content.Content,
		Markdown: content.Markdown,
		Ip:       clientIP}

	_, err := commentService.Update(content.Id, comment)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("Update success.")
}

func (c Comment) Destroy(ids [] int) revel.Result {
	if len(ids) <= 0 {
		return c.ResponseError(402, "参数不能为空")
	}

	for _, id := range ids {
		_, err := commentService.Destroy(id, models.Comments{})
		if err != nil {
			return c.ResponseError(500, err.Error())
		}
	}
	return c.ResponseSuccess("Delete success.")
}

func (c Comment) Spam(id int, status int) revel.Result {
	c.Validation.Required(id)
	c.Validation.Required(status)
	if c.Validation.HasErrors() {
		return c.RenderJSON(models.Res{Status: 501, Info: c.Validation.Errors})
	}

	comment := models.Comments{Status: status}
	_, err := commentService.Spam(id, comment)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("success.")
}
