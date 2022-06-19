package handler

import (
	"net/http"

	"github.com/LightAlykard/GoBackEnd-2/hw1/app/models/community"
	"github.com/LightAlykard/GoBackEnd-2/hw1/app/models/user"
)

type Router struct {
	*http.ServeMux
	us *user.Users
	cs *community.Communities
}

func NewRouter(us *user.Users, cs *community.Communities) *Router {
	return &Router{}
}
