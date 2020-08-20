package dto

import "mychat/internal/repo/dao"

//easyjson:json
type Chats []dao.Chat

type ChatReq struct {
	Name string `json:"name"`
	Users []uint32 `json:"users"`
}

type ChatMessagesReq struct {
	ID uint64 `json:"chat"`
}
