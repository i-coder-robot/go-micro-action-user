package handler

import (
	"context"
	"fmt"
	pb "github.com/i-coder-robot/go-micro-action-user/proto/role"
	"github.com/i-coder-robot/go-micro-action-user/service/repository"
)

type Role struct {
	Repo repository.Role
}

func (srv *Role) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	roles, err := srv.Repo.All(req)
	if err != nil {
		return err
	}
	res.Roles = roles
	return err
}

func (srv *Role) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	roles, err := srv.Repo.List(req.ListQuery)
	total, err := srv.Repo.Total(req.ListQuery)
	if err != nil {
		return err
	}
	res.Roles = roles
	res.Total = total
	return err
}

func (srv *Role) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	role, err := srv.Repo.Get(req.Role)
	if err != nil {
		return err
	}
	res.Role = role
	return err
}

func (srv *Role) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req.Role)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("添加角色失败 ")
	}
	res.Valid = true
	return err
}

func (srv *Role) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req.Role)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新角色失败")
	}
	res.Valid = valid
	return err
}

func (srv *Role) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	b, err := srv.Repo.Delete(req.Role)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除角色失败")
	}
	res.Valid = b
	return err
}
