package models

type Admin struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
