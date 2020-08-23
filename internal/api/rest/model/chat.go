package model

import "mychat/internal/repo/model"

//easyjson:json
type Chats []model.Chat

//easyjson:json
type ChatReq struct {
	Name string
	Users []uint32
}

//easyjson:json
type ChatMessagesReq struct {
	ID uint64
}
