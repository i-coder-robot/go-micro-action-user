package repository

import (
	"fmt"
	pb "github.com/i-coder-robot/go-micro-action-user/proto/user"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/util/log"
)

type User interface {
	Create(user *pb.User) (*pb.User, error)
	Exist(user *pb.User) bool
	Get(user *pb.User) (*pb.User, error)
	List(req *pb.ListQuery) ([]*pb.User, error)
	Total(req *pb.ListQuery) (int64, error)
	Update(user *pb.User) (bool, error)
	Delete(user *pb.User) (bool, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func (repo UserRepository) Exist(user *pb.User) bool {
	var count int
	if user.Username != "" {
		repo.DB.Model(&user).Where("username= ?", user.Username).Count(&count)
		if count > 0 {
			return true
		}
	}
	if user.Mobile != "" {
		repo.DB.Model(&user).Where("mobile = ?", user.Mobile).Count(&count)
		if count > 0 {
			return true
		}
	}
	if user.Email != "" {
		repo.DB.Model(&user).Where("email = ?", user.Email).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

func (repo UserRepository) List(req *pb.ListQuery) (users []*pb.User, err error) {
	db := repo.DB
	limit := req.Limit
	offset := req.Page * 10
	sort := req.Sort
	//TODO 查查这个 sort，Order()方法需要什么类型
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Total(req *pb.ListQuery) (total int64, err error) {
	users := []pb.User{}
	db := *repo.DB
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&users).Count(&total).Error; err != nil {
		log.Log(err)
		return total, err
	}
	return total, nil
}
func (repo *UserRepository) Get(user *pb.User) (*pb.User, error) {
	if err := repo.DB.Where(&user).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Create(user *pb.User) (*pb.User, error) {
	if exist := repo.Exist(user); exist == true {
		return user, fmt.Errorf("用户注册已存在")
	}
	err := repo.DB.Create(user).Error
	if err != nil {
		log.Log(err)
		return user, fmt.Errorf("用户注册失败")
	}
	return user, nil
}

func (repo *UserRepository) Update(user *pb.User) (bool, error) {
	if user.Id == "" {
		return false, fmt.Errorf("请传入更新 ID")
	}
	id := &pb.User{
		Id: user.Id,
	}
	err := repo.DB.Model(id).Update(user).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
func (repo *UserRepository) Delete(user *pb.User) (bool, error) {
	id := &pb.User{Id: user.Id}
	err := repo.DB.Delete(id).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
