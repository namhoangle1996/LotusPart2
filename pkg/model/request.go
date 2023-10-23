package model

type LoginRequest struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UploadFileRequest struct {
	File string `json:"file" form:"file" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
