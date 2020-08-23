package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"mychat/internal/api/rest/model"
	"mychat/internal/service"
	"net/http"
)

type ChatApi struct {
	service service.Service
}

func New(service service.Service) ChatApi {
	return ChatApi{service}
}

func (c *ChatApi) Router() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/users/add", c.addUser).Methods("POST")
	router.HandleFunc("/chats/add", c.createChat).Methods("POST")
	router.HandleFunc("/messages/add", c.sendMessage).Methods("POST")
	router.HandleFunc("/chats/get", c.getChatsByUserID).Methods("POST")
	router.HandleFunc("/messages/get", c.getMessagesByChatID).Methods("POST")

	return router
}

func (c *ChatApi) addUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user := model.User{}
	if err := json.Unmarshal(body, &user); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if user.Username == "" {
		writeError(w, http.StatusBadRequest, "Empty username")
		return
	}

	id, err := c.service.CreateUser(user.Username)
	if err != nil {
		if err == service.ErrUserExists {
			writeError(w, http.StatusConflict, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := json.Marshal(&model.Id{Id: id})
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if _, err := w.Write(resp); err != nil {
		log.Println("error write addUser response:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *ChatApi) createChat(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	chat := model.ChatReq{}
	if err := json.Unmarshal(body, &chat); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if chat.Name == "" {
		writeError(w, http.StatusBadRequest, "Empty name")
		return
	}

	if len(chat.Users) == 0 {
		writeError(w, http.StatusBadRequest, "Empty users")
		return
	}

	id, err := c.service.CreateChat(chat.Name, chat.Users)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := json.Marshal(&model.Id{Id: id})
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if _, err := w.Write(resp); err != nil {
		log.Println("error write createChat response:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *ChatApi) sendMessage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	message := model.MessageReq{}
	if err := json.Unmarshal(body, &message); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if message.Chat <= 0 {
		writeError(w, http.StatusBadRequest, "Wrong chat id")
		return
	}

	if message.Author <= 0 {
		writeError(w, http.StatusBadRequest, "Wrong user id")
		return
	}

	if message.Text == "" {
		writeError(w, http.StatusBadRequest, "Empty message")
		return
	}

	id, err := c.service.AddMessage(message.Chat, message.Author, message.Text)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := json.Marshal(&model.Id{Id: id})
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if _, err := w.Write(resp); err != nil {
		log.Println("error write sendMessage response:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *ChatApi) getChatsByUserID(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	user := model.UsersChatsReq{}
	if err := json.Unmarshal(body, &user); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if user.ID <= 0 {
		writeError(w, http.StatusBadRequest, "Wrong user id")
		return
	}

	chats, err := c.service.GetChatsByUserID(user.ID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := json.Marshal(model.Chats(chats))
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if _, err := w.Write(resp); err != nil {
		log.Println("error write getChatsByUserID response:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *ChatApi) getMessagesByChatID(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	chat := model.ChatMessagesReq{}
	if err := json.Unmarshal(body, &chat); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if chat.ID <= 0 {
		writeError(w, http.StatusBadRequest, "Wrong user id")
		return
	}

	messages, err := c.service.GetMessagesByChatID(chat.ID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := json.Marshal(model.Messages(messages))
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if _, err := w.Write(resp); err != nil {
		log.Println("error write GetMessagesByChatID response:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func writeError(w http.ResponseWriter, code int, body string) {
	w.WriteHeader(code)
	if _, err := w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, []byte(body)))); err != nil {
		log.Printf("error write response body: err: %v body: %s", err, body)
	}
}
