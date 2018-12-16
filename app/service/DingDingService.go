package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/revel/revel"
	"io/ioutil"
	"net/http"
)

var baseUrl = "https://oapi.dingtalk.com/robot/send?access_token=%s"

type TextMsg struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

type MarkdownMsg struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

type ResultMsg struct {
	Errmsg  string `json:"errmsg"`
	Errcode int    `json:"errcode"`
}

type DingdingService struct {
}

func (this *DingdingService) SendNotice(msg string) (bool, error) {
	token := revel.Config.StringDefault("ding.key", "")
	apiUrl := fmt.Sprintf(baseUrl, token)
	var textMsg MarkdownMsg
	textMsg.Msgtype = "markdown"
	textMsg.Markdown.Title = "Persimmon 有新评论"
	textMsg.Markdown.Text = msg
	textMsg.At.AtMobiles = make([]string, 0)
	textMsg.At.IsAtAll = false
	jsonStr, _ := json.Marshal(textMsg)
	request, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		revel.AppLog.Errorf("%s", err)
		return false, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	//revel.AppLog.Debugf("%s", string(body))

	if err != nil {
		revel.AppLog.Errorf("%s", err)

		return false, err
	}
	var resultMsg ResultMsg
	err = json.Unmarshal(body, &resultMsg)

	if resultMsg.Errcode != 0 {
		return false, errors.New(resultMsg.Errmsg)
	}

	return true, nil
}
