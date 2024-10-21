package models

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Patsurn  string `json:"patsurn"`
	Matsurn  string `json:"matsurn"`
	Addres   string `json:"addres"`
	Phone    string `json:"phone"`
	City     string `json:"city"`
	State    string `json:"state"`
	Username string `json:"username"`
	Password string `json:"password"`
	Rol      string `json:"rol"`
	Image    string `json:"image"`
}
