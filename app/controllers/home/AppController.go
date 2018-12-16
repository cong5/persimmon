package home

import (
	"fmt"
	"github.com/cong5/persimmon/app/models"
	"github.com/cong5/persimmon/app/utils"
	"github.com/revel/revel"
	"net/http"
)

const (
	PageSize  = 10
	FirstPage = 1
)

type AppController struct {
	*revel.Controller
}

type myRenderXML string

func (r myRenderXML) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusOK, "text/xml")
	resp.GetWriter().Write([]byte(r))
}

func (c AppController) AjaxSuccess(msg string) revel.Result {
	return c.RenderJSON(models.Res{Status: 200, Info: msg})
}

func (c AppController) AjaxError(code int, msg string) revel.Result {
	return c.RenderJSON(models.Res{Status: code, Info: msg})
}

func (c AppController) GetGlobalInfo() {
	allOption, err := optionService.GetAllOption(false)
	if err == nil {
		for _, v := range allOption {
			c.ViewArgs[v.Name] = v.Value
		}
	} else {
		revel.AppLog.Errorf("GetAllOption failed: %s", err)
	}

	//nav
	navi, navErr := navigationService.GetNavigation(false)
	if navErr != nil {
		revel.AppLog.Errorf("GetNavigation failed: %s", err)
	}
	c.ViewArgs["navigations"] = navi
}

func (c AppController) ViewAssign(posts []models.Posts, total int, limit int, page int) {
	c.GetGlobalInfo()

	posts = utils.SubstrContent(posts)
	totalPage := utils.GetTotalPage(total, limit)

	prefix := "page"
	if tagName := c.Params.Route.Get("name"); tagName != "" {
		prefix = fmt.Sprintf("tag/%s", tagName)
	}

	c.ViewArgs["posts"] = posts
	c.ViewArgs["page"] = page
	c.ViewArgs["total"] = total
	c.ViewArgs["totalPage"] = totalPage
	c.ViewArgs["showPage"] = postService.ShowPage(prefix, page, totalPage)
}
