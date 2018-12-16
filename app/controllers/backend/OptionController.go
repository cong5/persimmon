package backend

import (
	"github.com/cong5/persimmon/app/models"
	"github.com/revel/revel"
)

type Option struct {
	BaseController
}

func (c Option) Index(rows int, page int) revel.Result {
	lists, err := optionService.GetListPaging(rows, page)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(models.Res{Status: 200, List: lists})
}

func (c Option) Show(id int) revel.Result {
	option, err := optionService.GetOptionById(id)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(models.Res{Status: 200, Item: option})
}

func (c Option) Store(content *models.Options) revel.Result {
	//Validation
	c.Validation.Required(content.Title)
	c.Validation.Required(content.Name)
	if c.Validation.HasErrors() {
		return c.RenderJSON(models.Res{Status: 501, Info: c.Validation.Errors})
	}

	//save
	option := models.Options{Title: content.Title,
		Name: content.Name,
		Value: content.Value,
		Group: content.Group,
		Remark: content.Remark,
		Status: content.Status,
		DataType: content.DataType}

	_, err := optionService.Save(option)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("add success.")
}

func (c Option) Update(content *models.Options) revel.Result {

	//save
	option := models.Options{Title: content.Title,
		Name: content.Name,
		Value: content.Value,
		Group: content.Group,
		Remark: content.Remark,
		Status: content.Status,
		DataType: content.DataType}

	_, err := optionService.Update(content.Id, option)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("Update success.")
}

func (c Option) Destroy(ids[] int) revel.Result {
	if len(ids) <= 0 {
		return c.ResponseError(402, "参数不能为空")
	}

	for _, id := range ids {
		_, err := optionService.Destroy(id, models.Options{})
		if err != nil {
			return c.ResponseError(500, err.Error())
		}
	}

	return c.ResponseSuccess("Delete success.")
}
