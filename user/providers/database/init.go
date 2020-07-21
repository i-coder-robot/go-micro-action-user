package database

import (
	"github.com/jinzhu/gorm"
	"github.com/lecex/core/env"
	"github.com/micro/go-micro/v2/util/log"
)

var (
	DB *gorm.DB
)

type DBConf struct {
	Driver string
	// Host 主机地址
	Host string
	// Port 主机端口
	Port string
	// User 用户名
	User string
	// Password 密码
	Password string
	// DbName 数据库名称
	DbName string
	// Charset 数据库编码
	Charset string
}

func GetConnection() (*gorm.DB, error) {
	dbConf := &DBConf{
		Driver: env.Getenv("DB_DRIVER", "mysql"),
		// Host 主机地址
		Host: env.Getenv("DB_HOST", "127.0.0.1"),
		// Port 主机端口
		Port: env.Getenv("DB_PORT", "3306"),
		// User 用户名
		User: env.Getenv("DB_USER", "root"),
		// Password 密码
		Password: env.Getenv("DB_PASSWORD", "123456"),
		// DbName 数据库名称
		DbName: env.Getenv("DB_NAME", "user"),
		// Charset 数据库编码
		Charset: env.Getenv("DB_CHARSET", "utf8"),
	}

	db, err := gorm.Open("mysql", dbConf.User+":"+dbConf.Password+"@tcp("+dbConf.Host+":"+dbConf.Port+")/"+dbConf.DbName+"?charset=utf8")
	if err != nil {
		panic("连接数据库失败")
	}
	db.SingularTable(true)
	//defer db.Close()
	if err != nil {
		panic(err)
		return nil, err
	}
	return db, nil
}

func init() {
	var err error
	DB, err = GetConnection()
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}

}
