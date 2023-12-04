package model

type Credential struct {
	Id       string `json:"id" xml:"id" form:"id" query:"id"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
}
