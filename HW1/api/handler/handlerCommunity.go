package handler

import (
	"net/http"

	"github.com/LightAlykard/GoBackEnd-2/hw1/app/models/community"
	"github.com/LightAlykard/GoBackEnd-2/hw1/app/models/user"
	"github.com/google/uuid"
)

type Community struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	CommUs []string  `json:"communities"`
}

func (rt *Router) CreateComm(w http.ResponseWriter, r *http.Request) {
	// TODO: парсинг жсона
	_ = community.Community{
		Name:  "Name from json",
		Users: []user.User{},
		// TODO: вызов Create CommunityInterface
	}
}

func (rt *Router) AddUser(w http.ResponseWriter, r *http.Request) {
	// TODO: AddCommunity из пакета user
	// TODO: AddUser из пакета community
}

func (rt *Router) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO: DeleteCommunity из пакета user
	// TODO: DeleteUser из пакета community
}

func (rt *Router) Search(w http.ResponseWriter, r *http.Request) {
	// TODO: SearchUser через интерфейс CommInterface
}
