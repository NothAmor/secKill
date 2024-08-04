package proto

type UserInfo struct {
}

type (
	CreateUserReq struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	CreateUserResp struct {
		Id int64 `json:"id"`
	}
)
