package backend

import (
	"github.com/cong5/persimmon/app/models"
	"github.com/revel/revel"
	"strconv"
	"strings"
)

type BaseController struct {
	*revel.Controller
}

func (c BaseController) ResponseSuccess(msg string) revel.Result {
	return c.RenderJSON(models.Res{Status: 200, Info: msg})
}

func (c BaseController) ResponseError(code int, msg string) revel.Result {
	return c.RenderJSON(models.Res{Status: code, Info: msg})
}

func (c BaseController) GetSession(key string) string {
	v, ok := c.Session[key]
	if !ok {
		v = ""
	}
	return v.(string)
}

func (c BaseController) SetSession(userInfo *models.Users) {
	c.Session["UserId"] = strconv.Itoa(userInfo.Id)
	c.Session["Email"] = string(userInfo.Email)
	c.Session["Username"] = string(userInfo.Name)
}

func (c BaseController) ClearSession() {
	delete(c.Session, "UserId")
	delete(c.Session, "Email")
	delete(c.Session, "Username")
}

func (c BaseController) Input(key string) string {
	values := c.Params.Values
	if values == nil {
		return ""
	}

	valStr := values[key]
	if len(valStr) == 0 {
		return ""
	}
	return strings.Join(valStr, "")
}

func (c BaseController) RequestData() map[string]string {
	values := c.Params.Values
	request := make(map[string]string, len(values))
	for key, value := range values {
		request[key] = strings.Join(value, "")
	}
	return request
}

func (c BaseController) GetDataByJSON() map[string]interface{} {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	return jsonData
}
