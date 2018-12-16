package backend

import (
	"github.com/cong5/persimmon/app/models"
	"github.com/revel/revel"
)

type User struct {
	BaseController
}

func (c User) Index() revel.Result {
	email := c.GetSession("Email")
	user, err := userService.GetUserByEmail(email)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	return c.RenderJSON(models.Res{Status: 200, Item: user})
}

func (c User) Update(id int, oldPwd string, pwd string, rePwd string) revel.Result {
	//Validation
	c.Validation.Required(oldPwd)
	c.Validation.Required(pwd)
	c.Validation.Required(rePwd)
	if c.Validation.HasErrors() {
		return c.RenderJSON(models.Res{Status: 501, Info: c.Validation.Errors})
	}
	if pwd != rePwd {
		return c.ResponseError(501, "password and password_confirmation unequal.")
	}

	//Valid old password
	email := c.GetSession("Email")
	user, err := userService.GetUserByEmail(email)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}

	if authService.CheckPasswordHash(pwd, user.Password) {
		return c.ResponseError(500, "Old password error.")
	}

	_, upErr := userService.UpdatePassword(id, pwd)
	if upErr != nil {
		return c.ResponseError(500, "Update password failed.")
	}

	return c.ResponseSuccess("Update password success.")
}
