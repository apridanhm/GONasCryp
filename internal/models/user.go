package models

type User struct {
	ID           string
	Username     string
	PasswordHash string
	Role         string
}
