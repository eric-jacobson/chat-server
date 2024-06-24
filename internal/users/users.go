package users

import "time"

type UserReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserResp struct {
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginResp struct {
	AccessToken string `json:"access_token"`
}
