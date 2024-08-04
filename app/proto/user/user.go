package user

import "time"

type UserInfo struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserReq 创建用户请求
type (
	CreateUserReq struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	CreateUserResp struct {
		Id int64 `json:"id"`
	}
)

// UserLoginReq 登录请求
type (
	UserLoginReq struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	UserLoginResp struct {
		Token string `json:"token"`
	}
)
