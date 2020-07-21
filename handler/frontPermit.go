package handler

import (
	"context"
	"fmt"
	"github.com/go-playground/locales/vai_Latn"
	pb "github.com/i-coder-robot/go-micro-action-user/user/proto/frontPermit"
	"github.com/i-coder-robot/go-micro-action-user/user/service/repository"
)

type FrontPermit struct {
	Repo repository.FrontPermit
}

func (srv *FrontPermit) UpdateOrCreate(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	p := pb.FrontPermit{}
	p.App = req.FrontPermit.App
	p.Service = req.FrontPermit.Service
	p.Method = req.FrontPermit.Method

	frontPermit, err := srv.Repo.Get(&p)
	if frontPermit != nil {
		_, err = srv.Repo.Create(req.FrontPermit)
	} else {
		req.FrontPermit.Id = frontPermit.Id
		_, err = srv.Repo.Update(req.FrontPermit)
	}
	return err
}

func (srv *FrontPermit) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	frontPermits, err := srv.Repo.All(req)
	if err != nil {
		return err
	}
	res.FrontPermits = frontPermits
	return err
}

func (srv *FrontPermit) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	frontPermits, err := srv.Repo.List(req.ListQuery)
	total, err := srv.Repo.Total(req.ListQuery)
	if err != nil {
		return err
	}
	res.FrontPermits = frontPermits
	res.Total = total
	return err
}

func (srv *FrontPermit) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	frontPermit, err := srv.Repo.Get(req.FrontPermit)
	if err != nil {
		return err
	}
	res.FrontPermit = frontPermit
	return err
}

func (srv *FrontPermit) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req.FrontPermit)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("添加前端权限失败")
	}
	res.Valid = true
	return err
}
func (srv *FrontPermit) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req.FrontPermit)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新前端权限失败")
	}
	res.Valid = valid
	return err
}

func (srv *FrontPermit) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req.FrontPermit)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除前端权限失败")
	}
	res.Valid = valid
	return err
}
