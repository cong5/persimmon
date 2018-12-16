package models

type Res struct {
	Id     interface{} `json:"id"`
	Status int         `json:"status"`
	Info   interface{} `json:"info"`
	List   interface{} `json:"list"`
	Item   interface{} `json:"item"`
}
