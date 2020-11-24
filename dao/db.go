package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

func InitMySQL() (err error) {
	// 连接MySQL
	DB, err = gorm.Open("mysql", "root:123456@(10.0.6.194)/bubble?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("Connect mysql failed, err: %v\n", err)
		return err
	}
	// 测试MySQL连通性
	err = DB.DB().Ping()
	return
}