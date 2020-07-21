package handler

import (
	"context"
	"fmt"
	pb "github.com/i-coder-robot/go-micro-action/user/proto/permission"
	"github.com/i-coder-robot/go-micro-action/user/service/repository"
)

type Permission struct {
	Repo repository.Permission
}

func (srv *Permission) UpdateOrCreate(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	p := pb.Permission{}
	p.Service = req.Permission.Service
	p.Method = req.Permission.Method
	permission, err := srv.Repo.Get(&p)
	if permission != nil {
		_, err = srv.Repo.Create(req.Permission)
	} else {
		req.Permission.Id = permission.Id
		_, err = srv.Repo.Update(req.Permission)
	}
	return err
}

func (srv *Permission) Sync(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	for _, p := range req.Permissions {
		req.Permission = p
		err = srv.UpdateOrCreate(ctx, req, res)
		if err != nil {
			return err
		}
	}
	return nil
}

func (srv *Permission) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	permissions, err := srv.Repo.All(req)
	if err != nil {
		return err
	}
	res.Permissions = permissions
	return err
}

func (srv *Permission) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	permissions, err := srv.Repo.List(req.ListQuery)
	total, err := srv.Repo.Total(req.ListQuery)
	if err != nil {
		return err
	}
	res.Permissions = permissions
	res.Total = total
	return err
}

func (srv *Permission) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	permission, err := srv.Repo.Get(req.Permission)
	if err != nil {
		return err
	}
	res.Permission = permission
	return err
}

func (srv *Permission) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req.Permission)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("添加权限失败")
	}
	res.Valid = true
	return err
}

func (srv *Permission) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req.Permission)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新权限失败")
	}
	res.Valid = valid
	return err
}

func (srv *Permission) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req.Permission)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除权限失败")
	}
	res.Valid = valid
	return err
}
