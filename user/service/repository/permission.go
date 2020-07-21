package repository

import (
	"fmt"
	"github.com/go-log/log"
	pb "github.com/i-coder-robot/go-micro-action/user/proto/permission"
	"github.com/jinzhu/gorm"
)

type Permission interface {
	Create(permission *pb.Permission) (*pb.Permission, error)
	Delete(permission *pb.Permission) (bool, error)
	Update(permission *pb.Permission) (bool, error)
	Get(permission *pb.Permission) (*pb.Permission, error)
	All(req *pb.Request) ([]*pb.Permission, error)
	List(req *pb.ListQuery) ([]*pb.Permission, error)
	Total(req *pb.ListQuery) (int64, error)
}
type PermissionRepository struct {
	DB *gorm.DB
}

func (repo *PermissionRepository) All(req *pb.Request) (permissions []*pb.Permission, err error) {
	if err := repo.DB.Find(&permissions).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return permissions, nil
}
func (repo *PermissionRepository) List(req *pb.ListQuery) (permissions []*pb.Permission, err error) {
	db := repo.DB
	limit := req.Limit
	offset := req.Page * 10
	sort := req.Sort
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&permissions).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return permissions, nil
}

func (repo *PermissionRepository) Total(req *pb.ListQuery) (total int64, err error) {
	permissions := pb.Permission{}
	db := repo.DB
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&permissions).Count(&total).Error; err != nil {
		log.Log(err)
		return total, err
	}
	return total, nil
}

func (repo *PermissionRepository) Get(permission *pb.Permission) (*pb.Permission, error) {
	if err := repo.DB.Where(&permission).Find(&permission).Error; err != nil {
		return nil, err
	}
	return permission, nil
}

func (repo *PermissionRepository) Create(p *pb.Permission) (*pb.Permission, error) {
	err := repo.DB.Create(p).Error
	if err != nil {
		log.Log(err)
		return p, fmt.Errorf("添加权限失败")
	}
	return p, nil
}

func (repo *PermissionRepository) Update(p *pb.Permission) (bool, error) {
	if p.Id == 0 {
		return false, fmt.Errorf("请传入更新 ID")
	}
	id := &pb.Permission{Id: p.Id}
	err := repo.DB.Model(id).Updates(p).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

func (repo *PermissionRepository) Delete(p *pb.Permission) (bool, error) {
	err := repo.DB.Delete(p).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
