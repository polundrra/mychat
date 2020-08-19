package model

import "mychat/internal/repo/model"

//easyjson:json
type Chats []model.Chat

type ChatReq struct {
	Name string `json:"name"`
	Users []uint32 `json:"users"`
}

type ChatMessagesReq struct {
	ID uint64 `json:"chat"`
}
