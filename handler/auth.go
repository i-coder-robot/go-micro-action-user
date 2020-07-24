package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-log/log"
	pb "github.com/i-coder-robot/go-micro-action-user/proto/auth"
	userPb "github.com/i-coder-robot/go-micro-action-user/proto/user"
	"github.com/i-coder-robot/go-micro-action-user/service"
	"github.com/i-coder-robot/go-micro-action-user/service/repository"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	TokenService service.Authable
	Repo         repository.User
}

func (srv *Auth) AuthById(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	user, err := srv.Repo.Get(&userPb.User{
		Id: req.User.Id,
	})
	if err != nil {
		log.Log(err)
		return err
	}
	if user != nil {
		req.User.Id = user.Id
		t, err := srv.TokenService.Encode(req.User)
		if err != nil {
			log.Log(err)
			return err
		}
		res.Token = t
	}
	return nil
}

func (srv *Auth) Auth(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	fmt.Println("Auth.Auth")
	user, err := srv.Repo.Get(&userPb.User{
		Id:       req.User.Id,
		Username: req.User.Username,
		Mobile:   req.User.Mobile,
		Email:    req.User.Email,
	})
	if err != nil {
		return errors.New("用户不存在")
	}
	if user != nil {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.User.Password)); err != nil {
			log.Log(err)
			return errors.New("密码错误")
		}
		req.User.Id = user.Id
		req.User.Username = user.Username
		req.User.Email = user.Email
		req.User.Mobile = user.Mobile
		t, err := srv.TokenService.Encode(req.User)
		if err != nil {
			log.Log(err)
			return err
		}
		res.Token = t
	}
	return nil
}

func (srv *Auth) ValidateToken(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	if req.Token == "" {
		return errors.New("请输入 Token")
	}
	claims, err := srv.TokenService.Decode(req.Token)
	if err != nil {
		log.Log(err)
		return err
	}
	if claims.User.Id == "" {
		return errors.New("无效用户")
	}
	res.User = claims.User
	res.Valid = true
	return err
}
