package handler

import (
	userPB "github.com/i-coder-robot/go-micro-action-user/proto/user"
	"github.com/i-coder-robot/go-micro-action-user/providers/casbin"
	db "github.com/i-coder-robot/go-micro-action-user/providers/database"
	"github.com/i-coder-robot/go-micro-action-user/service"
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

// Auth 授权管理服务实现
func (srv *Handler) Auth() *Auth {
	return &Auth{&service.TokenService{}, &repository.UserRepository{db.DB}}
}

// FrontPermit 前端权限管理服务实现
func (srv *Handler) FrontPermit() *FrontPermit {
	return &FrontPermit{&repository.FrontPermitRepository{db.DB}}
}

// Permission 权限管理服务实现
func (srv *Handler) Permission() *Permission {
	return &Permission{&repository.PermissionRepository{db.DB}}
}

// Role 角色管理服务实现
func (srv *Handler) Role() *Role {
	return &Role{&repository.RoleRepository{db.DB}}
}

// Casbin 权限管理服务实现
func (srv *Handler) Casbin() *Casbin {
	return &Casbin{casbin.Enforcer}
}
