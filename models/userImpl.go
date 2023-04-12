package models

type UserRegisterInput struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Age      int    `json:"age" form:"age"`
}

type UserRegisterOutput struct {
	Base
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type UserLoginInput struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserLoginOutput struct {
	Token string `json:"token"`
}
