package models

type SearchPostReq struct {
	UserName   string `json:"username"`
	UserPasswd string `json:"userpasswd"`
}