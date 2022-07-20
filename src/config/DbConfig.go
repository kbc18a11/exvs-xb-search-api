package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig interface {
	DbConnectionInit() (*gorm.DB, error)
}

type DbConfigImp struct {
}

/*
DB接続初期化処理
*/
func (dbConfigImp *DbConfigImp) DbConnectionInit() (*gorm.DB, error) {
	dsn := "root:password@tcp(127.0.0.1:3306)/exvssearchxb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
