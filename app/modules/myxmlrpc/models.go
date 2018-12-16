package myxmlrpc

import "time"

type Base64 struct {
}

type BlogInfo struct {
	BlogId   string `xml:"blogid"`
	URL      string `xml:"url"`
	BlogName string `xml:"blogName"`
}

type Enclosure struct {
	Length int    `xml:"length"`
	Type   string `xml:"type"`
	URL    string `xml:"url"`
}

type Source struct {
	Name string `xml:"name"`
	URL  string `xml:"url"`
}

type Post struct {
	DateCreated time.Time `xml:"dateCreated"`
	Description string    `xml:"description"`
	Title       string    `xml:"title"`
	Categories  []string  `xml:"categories"`
	Enclosure   Enclosure `xml:"enclosure"`
	Link        string    `xml:"link"`
	Permalink   string    `xml:"permalink"`
	PostId      string    `xml:"postid"`
}

type CategoryInfo struct {
	Description string `xml:"description"`
	HtmlURL     string `xml:"htmlUrl"`
	RssURL      string `xml:"rssUrl"`
	Title       string `xml:"title"`
	CategoryId  string `xml:"categoryid"`
}

type MediaObject struct {
	Name string `xml:"name"`
	Type string `xml:"type"`
	Bits Base64 `xml:"bits"`
}

type MediaObjectUrl struct {
	URL string `xml:"url"`
}

type Params struct {
	Value interface{} `xml:"param"`
}
