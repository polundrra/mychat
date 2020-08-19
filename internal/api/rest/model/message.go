package model

import "mychat/internal/repo/model"

//easyjson:json
type MessageReq struct {
	Chat uint64 `json:"chat"`
	Author uint64 `json:"author"`
	Text string `json:"text"`
}

type Messages []model.Message
