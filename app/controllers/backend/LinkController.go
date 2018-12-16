package backend

import (
	"github.com/cong5/persimmon/app/models"
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/utils"
)

type Link struct {
	BaseController
}

func (c Link) Index(rows int, page int) revel.Result {
	lists, err := linkService.GetListPaging(rows, page)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(models.Res{Status: 200, List: lists})
}

func (c Link) Show(id int) revel.Result {
	link, err := linkService.GetLinkById(id,false)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(models.Res{Status: 200, Item: link})
}

func (c Link) Store(content *models.Links) revel.Result {

	//Validation
	c.Validation.Required(content.Name)
	c.Validation.Required(content.Url)
	if c.Validation.HasErrors() {
		return c.RenderJSON(models.Res{Status: 501, Info: c.Validation.Errors})
	}

	//save
	clientIP := utils.Ip2long(c.ClientIP)
	link := models.Links{Name: content.Name,
		Logo: content.Logo,
		Group: content.Group,
		Url: content.Url,
		Ip: clientIP}

	_, err := linkService.Save(link)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("add success.")
}

func (c Link) Update(content *models.Links) revel.Result {

	//save
	clientIP := utils.Ip2long(c.ClientIP)
	link := models.Links{Name: content.Name,
		Logo: content.Logo,
		Group: content.Group,
		Url: content.Url,
		Ip: clientIP}

	_, err := linkService.Update(content.Id, link)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("Update success.")
}

func (c Link) Destroy(ids[] int) revel.Result {
	if len(ids) <= 0 {
		return c.ResponseError(402, "参数不能为空")
	}

	for _, id := range ids {
		_, err := linkService.Destroy(id, models.Links{})
		if err != nil {
			return c.ResponseError(500, err.Error())
		}
	}

	return c.ResponseSuccess("Delete success.")
}
