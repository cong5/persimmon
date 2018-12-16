package models

import "encoding/xml"

type ServerRequest struct {
	Name   xml.Name `xml:"methodCall"`
	Method string   `xml:"methodName"`
	Rawxml string
}

type GetUsersBlogs struct {
	XMLName    xml.Name `xml:"methodCall"`
	Text       string   `xml:",chardata"`
	MethodName string   `xml:"methodName"`
	Params     struct {
		Text  string `xml:",chardata"`
		Param []struct {
			Text  string `xml:",chardata"`
			Value struct {
				Text   string `xml:",chardata"`
				String string `xml:"string"`
			} `xml:"value"`
		} `xml:"param"`
	} `xml:"params"`
}

type Member struct {
	Text  string `xml:",chardata"`
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

type LoginResponse struct {
	XMLName xml.Name `xml:"methodResponse"`
	Text    string   `xml:",chardata"`
	Ex      string   `xml:"ex,attr"`
	Params  struct {
		Text  string `xml:",chardata"`
		Param struct {
			Text  string `xml:",chardata"`
			Value struct {
				Text  string `xml:",chardata"`
				Array struct {
					Text string `xml:",chardata"`
					Data struct {
						Text  string `xml:",chardata"`
						Value struct {
							Text   string `xml:",chardata"`
							Struct struct {
								Text   string `xml:",chardata"`
								Member *[]Member `xml:"member"`
							} `xml:"struct"`
						} `xml:"value"`
					} `xml:"data"`
				} `xml:"array"`
			} `xml:"value"`
		} `xml:"param"`
	} `xml:"params"`
}
