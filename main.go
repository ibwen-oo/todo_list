package main

import (
	"bubble/dao"
	"bubble/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)


type Todo struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func main() {
	// 解决windows下gin项目在console输出的日志乱码问题
	gin.DisableConsoleColor()

	// 初始化MySQL连接
	if err := dao.InitMySQL(); err != nil {
		log.Printf("Init DB failed, err: %v", err)
		return
	}
	// 程序退出,关闭MySQL连接
	defer dao.DB.Close()
	// 模型绑定
	dao.DB.AutoMigrate(&Todo{})

	// 设置路由
	r := routers.SetupRouters()

	r.Run("0.0.0.0:9090")
}
