package main

import (
	"log"
	"mychat/internal/api/rest"
	"mychat/internal/repo"
	"mychat/internal/service"
	"net/http"
)

func main() {
	repo, err := repo.New(repo.Opts{
		Host: "db",
		Port: 5432,
		Database: "mychat",
		User: "polina",
		Password: "super_secret",
	})
	if err != nil {
		log.Fatal(err)
	}

	service := service.New(repo)
	api := rest.New(service)

	if err := http.ListenAndServe(":9000", api.Router()); err != nil {
		log.Fatal(err)
	}
}

//curl --request POST --data '{"username": "alex"}' http://localhost:9000/users/add
//curl --request POST --data '{"chat": 21, "author": 4, "text": "what's up"}' http://localhost:9000/messages/add
//curl --request POST --data '{"name": "chat_1", "users": [1, "2"]}' http://localhost:9000/chats/add
//curl --request POST --data '{"user": 1}' http://localh ost:9000/chats/get
//curl --request POST --data '{"chat": "10"}' http://localhost:9000/messages/get