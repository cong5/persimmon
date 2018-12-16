package backend

import (
	"github.com/cong5/persimmon/app/models"
	"github.com/revel/revel"
)

type Settings struct {
	BaseController
}

func (c Settings) Index() revel.Result {
	options, err := optionService.GetAllOption(false)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(models.Res{Status: 200, List: options})
}

func (c Settings) Update() revel.Result {
	data := c.RequestData()
	for key, value := range data {
		optionService.UpdateByName(key, value)
	}
	return c.ResponseSuccess("update success.")
}
