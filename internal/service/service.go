package service

import (
	"mychat/internal/repo"
	"mychat/internal/repo/dao"
)

type Service interface {
	CreateUser(username string) (uint64, error)
	CreateChat(name string, users []uint32) (uint64, error)
	AddMessage(chat uint64, author uint64, text string) (uint64, error)
	GetChatsByUserID(userID uint32) ([]dao.Chat, error)
	GetMessagesByChatID(chatID uint64) ([]dao.Message, error)
}

func New(repo repo.ChatRepo) Service {
	return chatService{repo: repo}
}
