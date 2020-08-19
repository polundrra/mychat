package model

//easyjson:json
type User struct {
	Username string `json:"username"`
}

type UsersChatsReq struct {
	ID uint32 `json:"user"`
}
