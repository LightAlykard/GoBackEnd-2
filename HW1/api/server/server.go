package server

import (
	"context"
	"net/http"
	"time"

	"github.com/LightAlykard/GoBackEnd-2/hw1/app/models/community"
	"github.com/LightAlykard/GoBackEnd-2/hw1/app/models/user"
)

type Server struct {
	srv http.Server
	us  *user.Users
	cs  *community.Communities
}

func NewServer(addr string, h http.Handler) *Server {
	s := &Server{}

	s.srv = http.Server{
		Addr:              addr,
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
	return s
}

func (s *Server) Start(us *user.Users, cs *community.Communities) {
	s.us = us
	s.cs = cs
	go s.srv.ListenAndServe()
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	s.srv.Shutdown(ctx)
	cancel()
}
