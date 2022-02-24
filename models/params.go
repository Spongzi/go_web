package models

type ParamsUser struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required,gte=8,lte=16"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"`
}

type ParamsLoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
