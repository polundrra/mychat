package service

import (
	"mychat/internal/repo"
	"mychat/internal/repo/model"
)

type chatService struct {
	repo repo.ChatRepo
}

func (c chatService) CreateUser(username string) (uint64, error) {
	user, err := c.repo.GetUserByUsername(username)
	if err != nil {
		return 0, err
	}
	if user != nil {
		return 0, ErrUserExists
	}
	return c.repo.AddUser(username)
}

func (c chatService) CreateChat(name string, users []uint32) (uint64, error) {
	chat, err := c.repo.GetChatByName(name)
	if err != nil {
		return 0, err
	}
	if chat != nil {
		if isExists := equal(users, chat.Users); isExists == true {
			return 0, ErrChatExists
		}
	}
	return c.repo.CreateChat(name, users)
}

func equal(a, b []uint32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func (c chatService) AddMessage(chat uint64, author uint64, text string) (uint64, error) {
	chats, err := c.repo.GetChatsByUserID(uint32(author))
	if err != nil {
		return 0, err
	}
	for _, v := range chats {
		if v.ID == chat {
			return c.repo.AddMessage(chat, author, text)
		}
	}
	return 0, ErrChatNotFound
}

func (c chatService) GetChatsByUserID(userID uint32) ([]model.Chat, error) {
	user, err := c.repo.GetUserByUserID(uint64(userID))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return c.repo.GetChatsByUserID(userID)
}

func (c chatService) GetMessagesByChatID(chatID uint64) ([]model.Message, error) {
	chat, err := c.repo.GetChatByChatID(chatID)
	if err != nil {
		return nil, err
	}
	if chat == nil {
		return nil, ErrChatNotFound
	}
	return c.repo.GetMessagesByChatID(chatID)
}






