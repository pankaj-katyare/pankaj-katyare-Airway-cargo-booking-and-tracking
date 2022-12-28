package model

type Account struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
	Roles  string `json:"roles"`
	City   string `json:"city"`
}
