package dto

import "mychat/internal/repo/dao"

//easyjson:json
type MessageReq struct {
	Chat uint64 `json:"chat"`
	Author uint64 `json:"author"`
	Text string `json:"text"`
}

type Messages []dao.Message
