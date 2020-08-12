package middleware

import (
	"context"
	"errors"
	"fmt"
	client "github.com/i-coder-robot/go-micro-action-core/client"
	authPb "github.com/i-coder-robot/go-micro-action-user/proto/auth"
	casbinPb "github.com/i-coder-robot/go-micro-action-user/proto/casbin"
	PB "github.com/i-coder-robot/go-micro-action-user/proto/permission"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
)
// Handler 处理器
// 包含一些高阶函数比如中间件常用的 token 验证等
type Handler struct {
	UserService string
	Permissions []*PB.Permission
	Service     string
}

func (srv *Handler) Wrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) (err error) {
		if srv.IsAuth(req) {
			meta, ok := metadata.FromContext(ctx)
			if !ok {
				return errors.New("no auth meta-data found in request")
			}
			if token, ok := meta["X-Csrf-Token"]; ok {
				// Note this is now uppercase (not entirely sure why this is...)
				// token := strings.Split(meta["authorization"], "Bearer ")[1]
				// Auth here
				authRes := &authPb.Response{}
				err := client.Call(ctx, srv.UserService, "Auth.ValidateToken", &authPb.Request{
					Token: token,
				}, authRes)
				if err != nil || authRes.Valid == false {
					fmt.Println("校验token失败")
					return err
				}
				// 设置用户 id
				meta["Userid"] = authRes.User.Id
				meta["Username"] = authRes.User.Username
				fmt.Println("从token中解析出用户名:"+authRes.User.Username)
				meta["Email"] = authRes.User.Email
				meta["Mobile"] = authRes.User.Mobile
				meta["Service"] = srv.Service
				meta["Method"] = req.Method()
				ctx = metadata.NewContext(ctx, meta)
				if srv.IsPolicy(req) {
					// 通过 meta user_id 验证权限
					casbinRes := &authPb.Response{}
					err := client.Call(ctx, srv.UserService, "Casbin.Validate", &casbinPb.Request{}, casbinRes)
					if err != nil || casbinRes.Valid == false {
						return errors.New("权限拒绝")
					}
				}
			} else {
				return errors.New("Empty Authorization")
			}
		}
		err = fn(ctx, req, resp)
		return err
	}
}

// IsAuth 检测用户授权
func (srv *Handler) IsAuth(req server.Request) bool {
	fmt.Println("IsAuth。。。。")
	fmt.Println(srv.Permissions)
	for _, p := range srv.Permissions {
		if p.Method == req.Method() && p.Auth == true {
			srv.Service = p.Service
			return true
		}
	}
	return false
}

// IsPolicy 检查用户权限
func (srv *Handler) IsPolicy(req server.Request) bool {
	fmt.Println("IsPolicy.....")
	fmt.Println(srv.Permissions)
	for _, p := range srv.Permissions {
		if p.Method == req.Method() && p.Policy == true {
			return true
		}
	}
	return false
}
