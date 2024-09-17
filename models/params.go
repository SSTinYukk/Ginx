package models

type ParamSignUp struct {
	Username   string `json:"username,omitempty" binding:"required"`
	Password   string `json:"password,omitempty" binding:"required"`
	RePassword string `json:"re_password,omitempty" binding:"required,eqfield=Password"`
	Email      string `json:"email,omitempty" binding:"required"`
	Gender     string `json:"gender,omitempty" binding:"required"`
}
