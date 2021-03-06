package repo

import (
	"github.com/jackc/pgx"
	"mychat/internal/repo/model"
)

//ChatRepo is an abstraction over persistence storage
type ChatRepo interface {
	CreateChat(name string, users []uint32) (uint64, error)
	AddUser(username string) (uint64, error)
	AddMessage(chat uint64, author uint64, text string) (uint64, error)
	GetChatsByUserID(userID uint32) ([]model.Chat, error)
	GetMessagesByChatID(chatID uint64) ([]model.Message, error)

	GetUserByUsername(username string) (*model.User, error)
	GetChatByName(name string) (*model.Chat, error)
	GetUserByUserID(id uint64) (*model.User, error)
	GetChatByChatID(id uint64) (*model.Chat, error)
}

type Opts struct {
	Host string
	Port uint16
	Database string
	User string
	Password string
}

func New(opts Opts) (ChatRepo, error) {
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{Host: opts.Host, Port: opts.Port, Database: opts.Database, User: opts.User, Password: opts.Password},
	})
	if err != nil {
		return nil, err
	}
	repo := postgreRepo{pool}
	return &repo, nil
}
