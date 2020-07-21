package handler

import (
	"context"
	"fmt"
	pb "github.com/i-coder-robot/go-micro-action/user/proto/user"
	"github.com/i-coder-robot/go-micro-action/user/service/repository"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Repo repository.User
}

//func (srv *User )()  {
//
//}

func (srv *User) Exist(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	res.Valid = srv.Repo.Exist(req.User)
	return err
}

func (srv *User) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	user, err := srv.Repo.Get(req.User)
	if err != nil {
		return err
	}
	res.User = user
	return err
}

func (srv *User) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.User.Password = string(hashPassword)
	user, err := srv.Repo.Create(req.User)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("创建用户失败")
	}
	res.User = user
	res.Valid = true
	return err
}

func (srv *User) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	if req.User.Password != "" {
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.User.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		req.User.Password = string(hashedPass)
	}
	valid, err := srv.Repo.Update(req.User)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新用户失败")
	}
	res.Valid = valid
	return err
}

func (srv *User) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req.User)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除用户失败")
	}
	res.Valid = valid
	return err
}
