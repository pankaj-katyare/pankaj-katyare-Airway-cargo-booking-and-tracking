package model

type Account struct {
	Id          string `json:"id" `
	Name        string `json:"name"`
	CompanyName string `json:"company_name"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	Roles       string `json:"roles"`
	City        string `json:"city"`
	Password    string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenData struct {
	CustomerID string `json:"id"`
	Role       string `json:"role"`
	IsCustomer bool   `json:"is_customer"`
}
