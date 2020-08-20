package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"mychat/internal/api/rest"
	"mychat/internal/repo"
	"mychat/internal/service"
	"net/http"
)

type conf struct {
	ServerPort string
	RepoOpts repo.Opts
}

func main() {
	var conf conf
	if _, err := toml.DecodeFile("/etc/mychat/conf.toml", &conf); err != nil {
		log.Fatal(err)
	}
	repo, err := repo.New(conf.RepoOpts)
	if err != nil {
		log.Fatal(err)
	}

	service := service.New(repo)

	api := rest.New(service)

	if err := http.ListenAndServe(conf.ServerPort, api.Router()); err != nil {
		log.Fatal(err)
	}
}

//curl --request POST --data '{"username": "alex"}' http://localhost:9000/users/add
//curl --request POST --data '{"chat": 21, "author": 4, "text": "what's up"}' http://localhost:9000/messages/add
//curl --request POST --data '{"name": "chat_1", "users": [1, "2"]}' http://localhost:9000/chats/add
//curl --request POST --data '{"user": 1}' http://localh ost:9000/chats/get
//curl --request POST --data '{"chat": "10"}' http://localhost:9000/messages/get