package handler

import (
	authPB "github.com/i-coder-robot/go-micro-action-user/proto/auth"
	casbinPB "github.com/i-coder-robot/go-micro-action-user/proto/casbin"
	frontPermitPB "github.com/i-coder-robot/go-micro-action-user/proto/frontPermit"
	permissionPB "github.com/i-coder-robot/go-micro-action-user/proto/permission"
	rolePB "github.com/i-coder-robot/go-micro-action-user/proto/role"
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
	authPB.RegisterAuthHandler(srv.Server, srv.Auth())                       // token 服务实现
	frontPermitPB.RegisterFrontPermitsHandler(srv.Server, srv.FrontPermit()) // 前端权限服务实现
	permissionPB.RegisterPermissionsHandler(srv.Server, srv.Permission())    // 权限服务实现
	rolePB.RegisterRolesHandler(srv.Server, srv.Role())                      // 角色服务实现
	casbinPB.RegisterCasbinHandler(srv.Server, srv.Casbin())
}

func (srv *Handler) User() *User {
	return &User{Repo: &repository.UserRepository{
		DB: db.DB,
	}}
}

// Auth 授权管理服务实现
func (srv *Handler) Auth() *Auth {
	return &Auth{TokenService: &service.TokenService{},
		Repo: &repository.UserRepository{DB: db.DB}}
}

// FrontPermit 前端权限管理服务实现
func (srv *Handler) FrontPermit() *FrontPermit {
	return &FrontPermit{Repo:&repository.FrontPermitRepository{DB:db.DB}}
}

// Permission 权限管理服务实现
func (srv *Handler) Permission() *Permission {
	return &Permission{Repo:&repository.PermissionRepository{DB:db.DB}}
}

// Role 角色管理服务实现
func (srv *Handler) Role() *Role {
	return &Role{Repo:&repository.RoleRepository{DB:db.DB}}
}

// Casbin 权限管理服务实现
func (srv *Handler) Casbin() *Casbin {
	return &Casbin{Enforcer: casbin.Enforcer}
}
