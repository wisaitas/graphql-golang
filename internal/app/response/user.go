package response

import "github.com/wisaitas/graphql-golang/internal/app/model"

type UserResponse struct {
	BaseResponse
	User *model.User `json:"user"`
}
