package rest

import (
	"mychat/internal/service"
)

func New(service service.Service) ChatApi {
	return ChatApi{service}
}
