package repository

import (
	"fmt"
	"github.com/go-log/log"
	pb "github.com/i-coder-robot/go-micro-action-user/proto/frontPermit"
	"github.com/jinzhu/gorm"
)

type FrontPermit interface {
	Create(frontPermit *pb.FrontPermit) (*pb.FrontPermit, error)
	Delete(frontPermit FrontPermit) (bool, error)
	Update(frontPermit *pb.FrontPermit) (bool, error)
	Get(frontPermit FrontPermit) (*pb.FrontPermit, error)
	All(req *pb.Request) ([]*pb.FrontPermit, error)
	List(req *pb.ListQuery) ([]*pb.FrontPermit, error)
	Total(req *pb.ListQuery) (int64, error)
}
type FrontPermitRepository struct {
	DB *gorm.DB
}

func (repo *FrontPermitRepository) All(req *pb.Request) (frontPermits []*pb.FrontPermit, err error) {
	if err := repo.DB.Find(&frontPermits).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return frontPermits, nil
}

func (repo *FrontPermitRepository) List(req *pb.ListQuery) (frontPermits []*pb.FrontPermit, err error) {
	db := repo.DB
	limit := req.Limit
	offset := req.Page * 10
	sort := req.Sort
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Or(sort).Limit(limit).Offset(offset).Find(&frontPermits).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return frontPermits, nil
}

func (repo FrontPermitRepository) Total(req *pb.ListQuery) (total int64, err error) {
	frontPermits := []pb.FrontPermit{}
	db := repo.DB
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&frontPermits).Count(&total).Error; err != nil {
		log.Log(err)
		return total, err
	}
	return total, nil
}

func (repo *FrontPermitRepository) Get(frontPermit *pb.FrontPermit) (*pb.FrontPermit, error) {
	if err := repo.DB.Where(&frontPermit).Find(&frontPermit).Error; err != nil {
		return nil, err
	}
	return frontPermit, nil
}

func (repo PermissionRepository) Create(p *pb.FrontPermit) (*pb.FrontPermit, error) {
	err := repo.DB.Create(p).Error
	if err != nil {
		log.Log(err)
		return p, fmt.Errorf("添加前端权限失败")
	}
	return p, nil
}

func (repo *FrontPermitRepository) Update(p *pb.FrontPermit) (bool, error) {
	if p.Id == 0 {
		return false, fmt.Errorf("请传入更新 Id")
	}
	id := &pb.FrontPermit{Id: p.Id}
	err := repo.DB.Model(id).Updates(p).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

func (repo FrontPermitRepository) Delete(p *pb.FrontPermit) (bool, error) {
	err := repo.DB.Delete(p).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
