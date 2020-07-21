package repository

import (
	"fmt"
	pb "github.com/i-coder-robot/go-micro-action/user/proto/role"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/util/log"
)

type Role interface {
	Create(role *pb.Role) (*pb.Role, error)
	Delete(role *pb.Role) (bool, error)
	Update(role *pb.Role) (bool, error)
	Get(role *pb.Role) (*pb.Role, error)
	All(req *pb.Request) ([]*pb.Role, error)
	List(req *pb.ListQuery) ([]*pb.Role, error)
	Total(req *pb.ListQuery) (int64, error)
}
type RoleRepository struct {
	DB *gorm.DB
}

func (repo *RoleRepository) All(req *pb.Request) (roles []*pb.Role, err error) {
	if err := repo.DB.Find(&roles).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return roles, nil
}

func (repo *RoleRepository) List(req *pb.ListQuery) (roles []*pb.Role, err error) {
	db := repo.DB
	limit := req.Limit
	offset := req.Page * 10
	sort := req.Sort
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&roles).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return roles, nil
}

func (repo *RoleRepository) Total(req *pb.ListQuery) (total int64, err error) {
	roles := []pb.Role{}
	db := repo.DB
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&roles).Count(&total).Error; err != nil {
		log.Log(err)
		return total, err
	}
	return total, nil
}

func (repo *RoleRepository) Get(role *pb.Role) (*pb.Role, error) {
	if err := repo.DB.Where(&role).Find(&role).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return role, nil
}

func (repo RoleRepository) Create(r *pb.Role) (*pb.Role, error) {
	err := repo.DB.Create(r).Error
	if err != nil {
		log.Log(err)
		return r, fmt.Errorf("添加角色失败")
	}
	return r, nil
}

func (repo *RoleRepository) Update(r *pb.Role) (bool, error) {
	if r.Id == 0 {
		return false, fmt.Errorf("请输入更新ID")
	}
	id := &pb.Role{Id: r.Id}
	err := repo.DB.Model(id).Updates(r).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

func (repo *RoleRepository) Delete(r *pb.Role) (bool, error) {
	err := repo.DB.Delete(r).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
