package entity

type CustomerUser struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	City      string `json:"city"`
	Email     string `json:"email"`
	Passwords string `json:"passwords"`
}
