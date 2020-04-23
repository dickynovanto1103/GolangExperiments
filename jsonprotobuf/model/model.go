package model

type Student struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	IdCard IDCard `json:"idCard"`
}

type IDCard struct {
	Id int32 `json:"id"`
}