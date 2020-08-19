package service

type ChatError string

func (e ChatError) Error() string {
	return string(e)
}

const (
	ErrUserExists ChatError = "User already exists"
	ErrChatExists ChatError = "Chat already exists"
	ErrChatNotFound ChatError = "Chat not found"
	ErrUserNotFound ChatError = "User is not exist"
)
