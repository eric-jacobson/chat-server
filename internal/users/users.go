package users

import "time"

type CreateUserReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserResp struct {
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
