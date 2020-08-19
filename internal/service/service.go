package service

import (
	"mychat/internal/repo"
	"mychat/internal/repo/model"
)

type Service interface {
	CreateUser(username string) (uint64, error)
	CreateChat(name string, users []uint32) (uint64, error)
	AddMessage(chat uint64, author uint64, text string) (uint64, error)
	GetChatsByUserID(userID uint32) ([]model.Chat, error)
	GetMessagesByChatID(chatID uint64) ([]model.Message, error)
}

func New(repo repo.ChatRepo) Service {
	return chatService{repo: repo}
}
