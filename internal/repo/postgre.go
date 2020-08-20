package repo

import (
	"fmt"
	"github.com/jackc/pgx"
	"log"
	"mychat/internal/repo/dao"
	"strings"
	"time"
)

type postgreRepo struct {
	pool *pgx.ConnPool
}

func (p *postgreRepo) GetUserByUserID(id uint64) (*dao.User, error) {
	user := dao.User{}
	if err := p.pool.QueryRow("select * from users where id = $1", id).Scan(&user.ID, &user.Username, &user.CreatedAt); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (p *postgreRepo) GetChatByChatID(id uint64) (*dao.Chat, error) {
	chat := dao.Chat{}
	sql := "SELECT c.id, c.name, array(select user_id from users_chat where users_chat.chat_id = c.id), c.created_at " +
		"FROM chat c INNER JOIN users_chat u ON c.id = u.chat_id WHERE c.id = $1 LIMIT 1"
	if err := p.pool.QueryRow(sql, id).Scan(&chat.ID, &chat.Name, &chat.Users, &chat.CreatedAt); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &chat, nil
}

func (p *postgreRepo) CreateChat(name string, users []uint32) (uint64, error) {
	tx, err := p.pool.Begin()
	if err != nil {
		return 0, err
	}
	defer func(tx *pgx.Tx) {
		if err := tx.Rollback(); err != nil && err != pgx.ErrTxClosed {
			log.Printf("error rollback create chat %v", err)
		}
    }(tx)
	id := uint64(0)
	if err := tx.QueryRow("insert into chat(name) values ($1) returning id", name).Scan(&id); err != nil {
		log.Println("err insert into chat:", err)
		return 0, err
	}

	query, args := insertManyChatUsersQuery(users)
	if _, err := tx.Exec(query, append([]interface{}{id}, args...)...); err != nil {
		log.Println("err insert into users_chat:", err)
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("error commit CreateChat transaction %v", err)
		return 0, err
	}
	return id, nil
}

func insertManyChatUsersQuery(users []uint32) (string, []interface{}) {
	args := make([]interface{}, 0, len(users))
	buf := strings.Builder{}
	buf.WriteString("insert into users_chat(chat_id, user_id) values ")
	for i := range users {
		args = append(args, interface{}(users[i]))
		if i == len(users) - 1 {
			buf.WriteString(fmt.Sprintf("($1, $%d)", i + 2))
			break
		}
		buf.WriteString(fmt.Sprintf("($1, $%d),", i + 2))
	}

	return buf.String(), args
}

func (p *postgreRepo) GetChatByName(name string) (*dao.Chat, error) {
	chat := dao.Chat{}
	sql := "SELECT c.id, c.name, array(select user_id from users_chat where users_chat.chat_id = c.id), c.created_at " +
		"FROM chat c INNER JOIN users_chat u ON c.id = u.chat_id WHERE c.name = $1 LIMIT 1"
	err := p.pool.QueryRow(sql, name).Scan(&chat.ID, &chat.Name, &chat.Users, &chat.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &chat, nil
}

func (p *postgreRepo) AddUser(username string) (uint64, error) {
	id := uint64(0)
	sql := "insert into users(username) values ($1) returning id"
	if err := p.pool.QueryRow(sql, username).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (p *postgreRepo) GetUserByUsername(username string) (*dao.User, error) {
	user := dao.User{}
	err := p.pool.QueryRow("select * from users where username = $1", username).Scan(&user.ID, &user.Username, &user.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}



func (p *postgreRepo) AddMessage(chat uint64, author uint64, text string) (uint64, error) {
	id := uint64(0)
	sql := "insert into message(chat_id, user_id, body) values ($1, $2, $3) returning id"
	if err := p.pool.QueryRow(sql, chat, author, text).Scan(&id); err != nil {
		return 0, err
	}
	_, err := p.pool.Exec("update chat set recent_msg_at = $1 where id = $2", time.Now(), chat)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *postgreRepo) GetChatsByUserID(userID uint32) ([]dao.Chat, error) {
	sql := "SELECT c.id, c.name, array(select user_id from users_chat where users_chat.chat_id = c.id), c.created_at " +
		"FROM chat c INNER JOIN users_chat u ON c.id = u.chat_id WHERE u.user_id = $1 ORDER BY c.recent_msg_at DESC"
	chatsID := []dao.Chat{}
	chat := dao.Chat{}
	rows, err := p.pool.Query(sql, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&chat.ID, &chat.Name, &chat.Users, &chat.CreatedAt); err != nil {
			return nil, err
		}
		chatsID = append(chatsID, chat)
	}
	return chatsID, err
}

func (p *postgreRepo) GetMessagesByChatID(chatID uint64) ([]dao.Message, error) {
	sql := "SELECT * FROM message WHERE chat_id = $1 ORDER BY created_at ASC"
	messages := make([]dao.Message, 0)
	message := dao.Message{}
	rows, err := p.pool.Query(sql, chatID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&message.ID, &message.Chat, &message.Author, &message.Text, &message.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, err
}



