package entity

type User struct {
	ID           uint   `json:"id"`
	FirstName    string `json:"first-name"`
	LastName     string `json:"last-name"`
	Username     string `json:"user-name"`
	UserPassword string `json:"user-password"`
}
