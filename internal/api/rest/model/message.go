package model

import "mychat/internal/repo/model"

//easyjson:json
type MessageReq struct {
	Chat uint64
	Author uint64
	Text string
}

//easyjson:json
type Messages []model.Message
