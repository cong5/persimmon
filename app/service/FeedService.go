package service

import (
	"time"
	"github.com/cong5/persimmon/app/models"
	"github.com/cong5/persimmon/app/utils"
	"html/template"
)

type FeedService struct{}

func (this *FeedService) BuildFeed(domain string) (string, error) {
	allOption, optErr := optionService.GetAllOption(false)
	if optErr != nil {
		return "", optErr
	}
	optArr := make(map[string]string)
	for _, val := range allOption {
		optArr[val.Name] = val.Value
	}
	optArr["domain"] = domain

	posts, pErr := postService.SearchList(0, "", 15, 1, false)
	if pErr != nil {
		return "", pErr
	}

	return this.makeXml(optArr, posts), nil
}

func (this *FeedService) makeXml(optArr map[string]string, posts []models.Posts) string {
	lastBuildDate := time.Now().Format("Mon, 2 Jan 2006 15:04:05")
	begin := `<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns="http://www.w3.org/2005/Atom">`

	siteInfo := `<title>` + optArr["site_name"] + `</title>
        <description>` + optArr["description"] + `</description>
        <language>zh-CN</language>
        <updated>` + lastBuildDate + `</updated>
 		<link href="https://` + optArr["domain"] + `/"></link>
  		<link ref="self" href="https://` + optArr["domain"] + `/feed"></link>
		<copyright>Copyright (c) ` + optArr["site_name"] + `</copyright>`

	items := ""
	for _, val := range posts {
		user, _ := userService.GetUserByUid(val.UserId)
		if user.Name == "" {
			user.Name = optArr["site_name"]
		}
		pubDate := utils.Date("Mon, 2 Jan 2006 15:04:05", val.CreatedAt)
		items += `<entry>
            <title>` + val.Title + `</title>
			<link href="https://` + optArr["domain"] + "/post/" + val.Slug + `"  rel="alternate"></link>
    		<author>
      			<name>` + user.Name + `</name>
    		</author>
            <summary type="html"><![CDATA[` + template.HTMLEscapeString(val.Content) + `]]></summary>
            <guid>` + val.Slug + `</guid>
            <updated>` + pubDate + ` +0000</updated>
        </entry>`
	}

	end := `</feed>`
	html := begin + siteInfo + items + end
	return html
}
