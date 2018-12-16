package service

import (
	"bytes"
	"fmt"
	"github.com/cong5/persimmon/app/utils"
	"github.com/revel/revel"
	"html/template"
	"net/smtp"
	"strings"
)

type NotificationService struct{}

type Smilies struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

func (this *NotificationService) SendMail(subject string, body string) (bool, error) {
	host := revel.Config.StringDefault("email.host", "")
	port := revel.Config.StringDefault("email.port", "")
	account := revel.Config.StringDefault("email.account", "")
	password := revel.Config.StringDefault("email.password", "")
	to := revel.Config.StringDefault("email.to", "")

	auth := smtp.PlainAuth("", account, password, host)
	contentType := "Content-Type: text/html; charset=UTF-8"
	message := []byte("To: " + to + "\r\nFrom: " + account + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	serverAddr := fmt.Sprintf("%s:%s", host, port)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(serverAddr, auth, account, sendTo, message)

	if err != nil {
		revel.AppLog.Errorf("SendMail Error: %s", err.Error())
		return false, err
	}
	return true, nil
}

func (this *NotificationService) SendCommentNotice(postId int, commentId int, host string) (bool, error) {

	tplFile := "app/views/mails/comments.html"
	comment, cmErr := commentService.GetCommentById(commentId, false)
	if cmErr != nil {
		revel.AppLog.Errorf("GetCommentById Error: %s", cmErr)
		return false, cmErr
	}

	subject := fmt.Sprintf("“%s” 有新的评论", comment.Title)
	t, pErr := template.ParseFiles(tplFile)
	if pErr != nil {
		revel.AppLog.Errorf("Template ParseFiles Error: %s", pErr)
		return false, pErr
	}

	data := map[string]interface{}{
		"Title":     comment.Title,
		"Url":       comment.Url,
		"Name":      comment.Name,
		"Content":   template.HTML(comment.Content),
		"CreatedAt": utils.Date("2006-01-02 15:04:05", comment.CreatedAt),
		"Host":      host,
	}

	var tpl bytes.Buffer
	if exeErr := t.Execute(&tpl, data); exeErr != nil {
		revel.AppLog.Errorf("Template Execute Error: %s", exeErr)
		return false, exeErr
	}

	body := tpl.String()
	ret, err := this.SendMail(subject, body)
	if err != nil {
		return false, err
	}

	return ret, nil
}

func (this *NotificationService) SendDingNotice(postId int, commentId int, host string) (bool, error) {
	comment, cmErr := commentService.GetCommentById(commentId, false)
	if cmErr != nil {
		revel.AppLog.Errorf("GetCommentById Error: %s", cmErr)
		return false, cmErr
	}

	subject := fmt.Sprintf("#### “%s” 有新的评论 \n\n", comment.Title)
	createdAt := utils.Date("2006-01-02 15:04:05", comment.CreatedAt)
	commentContent := this.ParseSmilie(host, comment.Markdown)
	postUrl := fmt.Sprintf("[查看文章](https://%s/post/%s)", host, comment.Slug)
	commentStr := fmt.Sprintf("> 作者：%s \n\n > 地址：[%s](%s) \n\n > 时间：%s \n\n > 地址：%s \n\n > 内容：%s \n\n ", comment.Name, comment.Url, comment.Url, createdAt, postUrl, commentContent)
	text := subject + commentStr

	return dingdingService.SendNotice(text)
}

func (this *NotificationService) ParseSmilie(host string, content string) string {
	mapStr := [22]Smilies{
		{
			URL:  "/public/smilies/arrow.gif",
			Name: ":arrow:",
		},
		{
			URL:  "/public/smilies/biggrin.gif",
			Name: ":biggrin:",
		},
		{
			URL:  "/public/smilies/confused.gif",
			Name: ":confused:",
		},
		{
			URL:  "/public/smilies/cool.gif",
			Name: ":cool:",
		},
		{
			URL:  "/public/smilies/cry.gif",
			Name: ":cry:",
		},
		{
			URL:  "/public/smilies/eek.gif",
			Name: ":eek:",
		},
		{
			URL:  "/public/smilies/evil.gif",
			Name: ":evil:",
		},
		{
			URL:  "/public/smilies/exclaim.gif",
			Name: ":exclaim:",
		},
		{
			URL:  "/public/smilies/idea.gif",
			Name: ":idea:",
		},
		{
			URL:  "/public/smilies/lol.gif",
			Name: ":lol:",
		},
		{
			URL:  "/public/smilies/mad.gif",
			Name: ":mad:",
		},
		{
			URL:  "/public/smilies/mrgreen.gif",
			Name: ":mrgreen:",
		},
		{
			URL:  "/public/smilies/neutral.gif",
			Name: ":neutral:",
		},
		{
			URL:  "/public/smilies/question.gif",
			Name: ":question:",
		},
		{
			URL:  "/public/smilies/razz.gif",
			Name: ":razz:",
		},
		{
			URL:  "/public/smilies/redface.gif",
			Name: ":redface:",
		},
		{
			URL:  "/public/smilies/rolleyes.gif",
			Name: ":rolleyes:",
		},
		{
			URL:  "/public/smilies/sad.gif",
			Name: ":sad:",
		},
		{
			URL:  "/public/smilies/smile.gif",
			Name: ":smile:",
		},
		{
			URL:  "/public/smilies/surprised.gif",
			Name: ":surprised:",
		},
		{
			URL:  "/public/smilies/twisted.gif",
			Name: ":twisted:",
		},
		{
			URL:  "/public/smilies/wink.gif",
			Name: ":wink:",
		},
	}

	for _, v := range mapStr {
		url := fmt.Sprintf("![smilies](https://%s%s)", host, v.URL)
		content = strings.Replace(content, v.Name, url, -1)
	}

	return content
}
