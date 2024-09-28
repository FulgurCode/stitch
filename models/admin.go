package models

type Admin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type AdminPass struct {
	Username    string `json:"username" form:"username"`
	OldPassword string `json:"old-password" form:"old-password"`
	NewPassword string `json:"new-password" form:"new-password"`
}
