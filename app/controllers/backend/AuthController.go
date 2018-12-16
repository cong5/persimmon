package backend

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/models"
)

type Auth struct {
	BaseController
}

func (c Auth) Login(email string, password string) revel.Result {
	//Validation
	c.Validation.Required(email)
	c.Validation.Required(password)
	if c.Validation.HasErrors() {
		return c.RenderJSON(models.Res{Status: 400, Info: c.Validation.Errors})
	}

	//Login Auth
	userInfo, err := userService.GetUserByEmail(email)
	if err != nil {
		return c.ResponseError(500, err.Error())
	}
	AuthResult := authService.Login(password, userInfo)
	if AuthResult {
		c.SetSession(userInfo)
		userInfo.Password = "Do you want to know password? Sorry don't tell you."
		return c.RenderJSON(models.Res{Status: 200, Id: c.Session.ID(), Info: "Login success.", Item: userInfo})
	} else {
		return c.ResponseError(501, "Username or password error.")
	}

}

func (c Auth) Logout() revel.Result {
	c.ClearSession()
	return c.Redirect("/")
}
