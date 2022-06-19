package handler

import (
	"net/http"

	"github.com/LightAlykard/GoBackEnd-2/hw1/app/models/user"
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Communities []string  `json:"communities"`
}

func (rt *Router) Create(w http.ResponseWriter, r *http.Request) {
	bu := user.User{
		Name:            "Name from json",
		UserCommunities: []string{"test1", "test2"},
	}
	_, err := rt.us.Create(r.Context(), bu)
	if err != nil {
		http.Error(w, "error when creating", http.StatusInternalServerError)
		return
	}

}

func (rt *Router) AddCommunity(w http.ResponseWriter, r *http.Request) {
	// TODO: AddCommunity из пакета user
	// TODO: AddUser из пакета community
}

func (rt *Router) DeleteCommunity(w http.ResponseWriter, r *http.Request) {
	// TODO: DeleteCommunity из пакета user
	// TODO: DeleteUser из пакета community
}

func (rt *Router) SearchUser(w http.ResponseWriter, r *http.Request) {
	// TODO: SearchUser через интерфейс UserInterface
}
