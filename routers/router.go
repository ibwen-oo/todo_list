package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	// 创建路由引擎
	r := gin.Default()

	// 加载html模板文件
	r.LoadHTMLGlob("templates/*")

	// 配置静态文件
	r.Static("/static", "static")

	// 主页路由
	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		// 查看todo
		v1Group.GET("/todo", controller.GetTodoList)

		// 添加todo
		v1Group.POST("/todo", controller.AddTodo)

		// 删除某一条todo
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)

		// 修改某一条todo
		v1Group.PUT("/todo/:id", controller.ChangeTodo)
	}

	return r
}
