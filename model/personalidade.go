package model

type Personalidade struct{
	Id string `json:"id"`
	Nome string `json:"nome"`
	Descricao string `json:"descricao"`
}

var Personalidades []Personalidade