package backend

import (
	"github.com/cong5/persimmon/app/models"
	"github.com/revel/revel"
)

type Navigation struct {
	BaseController
}

func (c Navigation) Index() revel.Result {
	navigation, err := navigationService.GetNavigation(false)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(models.Res{Status: 200, List: navigation})
}

func (c Navigation) Update(nav string) revel.Result {
	//Validation
	c.Validation.Required(nav)
	if c.Validation.HasErrors() {
		return c.RenderJSON(models.Res{Status: 501, Info: c.Validation.Errors})
	}
	_, err := optionService.UpdateByName(navigationService.GetNavKey(), nav)
	if err != nil {
		return c.ResponseError(500, "update failed.")
	}

	return c.ResponseSuccess("update success.")
}
