package model

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type UserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	User    *User  `json:"user,omitempty"`
}
