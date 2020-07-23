package handler

import (
	userPB "github.com/i-coder-robot/go-micro-action-user/proto/user"
	db "github.com/i-coder-robot/go-micro-action-user/providers/database"
	"github.com/i-coder-robot/go-micro-action-user/service/repository"
	"github.com/micro/go-micro/v2/server"
)

type Handler struct {
	Server server.Server
}

func (srv *Handler) Register() {
	userPB.RegisterUsersHandler(srv.Server, srv.User())
}

func (srv *Handler) User() *User {
	return &User{Repo: &repository.UserRepository{
		DB: db.DB,
	}}
}
