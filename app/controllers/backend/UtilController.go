package backend

import (
	"github.com/cong5/persimmon/app/models"
	"github.com/revel/revel"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

type Utils struct {
	BaseController
}

const (
	baidu = "baidu"
	py    = "pinyin"
	//...
)

func (c Utils) Translate(words string) revel.Result {
	translateDriver := revel.Config.StringDefault("translate.driver", "pinyin")
	var res models.Res
	switch translateDriver {
	case baidu:
		if dst, err := baiduFanyiService.Fanyi(words); err == nil {
			res = models.Res{Status: 200, Item: dst}
		} else {
			res = models.Res{Status: 500, Info: err.Error()}
		}
	case py:
		newWords := pinyin.LazyPinyin(words, pinyin.NewArgs())
		wordsStrings := strings.Join(newWords, "")
		res = models.Res{Status: 200, Item: wordsStrings}
	default:
		res = models.Res{Status: 200, Item: ""}
	}

	return c.RenderJSON(res)
}
