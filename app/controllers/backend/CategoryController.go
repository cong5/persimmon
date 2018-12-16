package backend

import (
	"github.com/cong5/persimmon/app/models"
	"github.com/cong5/persimmon/app/utils"
	"github.com/revel/revel"
)

type Category struct {
	BaseController
}

func (c Category) Index(rows int, page int) revel.Result {
	lists, err := categoryService.GetListPaging(rows, page, false)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(models.Res{Status: 200, List: lists})
}

func (c Category) Show(id int) revel.Result {
	category, err := categoryService.GetCategoryById(id, false)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(models.Res{Status: 200, Item: category})
}

func (c Category) Store(content *models.Categorys) revel.Result {

	//Validation
	c.Validation.Required(content.Name)
	c.Validation.Required(content.Slug)
	if c.Validation.HasErrors() {
		return c.RenderJSON(models.Res{Status: 501, Info: c.Validation.Errors})
	}

	//save
	clientIP := utils.Ip2long(c.ClientIP)
	category := models.Categorys{Name: content.Name,
		Description: content.Description,
		Slug:        content.Slug,
		Pid:         content.Pid,
		Ip:          clientIP}

	_, err := categoryService.Save(category)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("add success.")
}

func (c Category) Update(content *models.Categorys) revel.Result {
	//save
	clientIP := utils.Ip2long(c.ClientIP)
	category := models.Categorys{Name: content.Name,
		Description: content.Description,
		Slug:        content.Slug,
		Pid:         content.Pid,
		Ip:          clientIP}

	_, err := categoryService.Update(content.Id, category)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.ResponseSuccess("Update success.")
}

func (c Category) Destroy(ids []int) revel.Result {
	_, err := categoryService.Destroy(ids, models.Categorys{})
	if err != nil {
		return c.ResponseError(500, "Delete failed.")
	}

	return c.ResponseSuccess("Delete success.")
}
