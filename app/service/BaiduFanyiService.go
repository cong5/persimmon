package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cong5/persimmon/app/models"
	"github.com/cong5/persimmon/app/utils"
	"github.com/revel/revel"
	"github.com/shurcooL/sanitized_anchor_name"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

type BaiduFanyiService struct{}

func (this BaiduFanyiService) Fanyi(words string) (string, error) {
	appId := revel.Config.StringDefault("baidu.translate.ak", "")
	secretKey := revel.Config.StringDefault("baidu.translate.sk", "")
	from := revel.Config.StringDefault("baidu.translate.from", "")
	to := revel.Config.StringDefault("baidu.translate.to", "")

	wordString, _ := this.Filter(words)
	salt := strconv.FormatInt(time.Now().Unix(), 10)
	signString := fmt.Sprintf("%s%s%s%s", appId, wordString, salt, secretKey)
	sign := utils.Md5(signString)

	//revel.AppLog.Info("%s", signString)
	//revel.AppLog.Info("%s", sign)

	baseUrl := "https://fanyi-api.baidu.com/api/trans/vip/translate?"
	query := map[string]string{
		"q":     url.QueryEscape(words),
		"from":  from,
		"to":    to,
		"appid": appId,
		"salt":  salt,
		"sign":  sign,
	}
	queryString := utils.Urlencode(query)
	resp, err := http.Get(baseUrl + queryString)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	baiduFanyi := models.BaiduFanyi{}
	jsonErr := json.Unmarshal(body, &baiduFanyi)
	if jsonErr != nil {
		return "", jsonErr
	}

	//Translation failed.
	if baiduFanyi.ErrorCode != "" {
		return "", errors.New(baiduFanyi.ErrorMsg)
	}

	dstWords := baiduFanyi.TransResult[0].Dst
	item := sanitized_anchor_name.Create(dstWords)
	return item, nil
}

func (this BaiduFanyiService) Filter(text string) (string, error) {
	if text == "" {
		return "", errors.New("Param empty.")
	}
	keyword := url.QueryEscape(text) //将关键字编码
	reg, err := regexp.Compile(`(%7E|%60|%21|%40|%23|%24|%25|%5E|%26|%27|%2A|%28|%29|%2B|%7C|%5C|%3D|\-|_|%5B|%5D|%7D|%7B|%3B|%22|%3A|%3F|%3E|%3C|%2C|\.|%2F|%A3%BF|%A1%B7|%A1%B6|%A1%A2|%A1%A3|%A3%AC|%7D|%A1%B0|%A3%BA|%A3%BB|%A1%AE|%A1%AF|%A1%B1|%A3%FC|%A3%BD|%A1%AA|%A3%A9|%A3%A8|%A1%AD|%A3%A4|%A1%A4|%A3%A1|%E3%80%82|%EF%BC%81|%EF%BC%8C|%EF%BC%9B|%EF%BC%9F|%EF%BC%9A|%E3%80%81|%E2%80%A6%E2%80%A6|%E2%80%9D|%E2%80%9C|%E2%80%98|%E2%80%99)+`)
	if err != nil {
		revel.AppLog.Errorf("FilterMark error: %s", err.Error())
		return text, err
	}
	keyword = reg.ReplaceAllString(keyword, "")
	keyword, unescErr := url.QueryUnescape(keyword) //将过滤后的关键字解码
	if unescErr != nil {
		revel.AppLog.Errorf("QueryUnescape error: %s", err.Error())
		return text, unescErr
	}
	return keyword, nil
}
