package controller

import (
	"bubble/dao"
	"bubble/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GetTodoList(c *gin.Context) {
	// 查询所有待办事项
	todoList, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func AddTodo(c *gin.Context) {
	// 获取用户输入(请求参数)
	var todo models.Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"Error": err.Error()})
		return
	}
	// 把用户输入保存到数据库中,并返回响应
	err := models.CreateTodo(&todo)
	if err != nil {
		log.Printf("Create Todo err: %v\n", err)
		c.JSON(http.StatusOK, gin.H{"Error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, "id不存在!")
		return
	}
	idx, _ := strconv.ParseInt(id, 10, 64)
	err := models.DeleteTodoByID(idx)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "Delete success!"})
	}
}

func ChangeTodo(c *gin.Context) {
	// 参数中取出待办事项的 id
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, "id 不存在!")
		return
	}
	// 从数据库中通过id检索待办事项
	var todo models.Todo
	idx, _ := strconv.ParseInt(id, 10, 64)
	err := models.ModifyTodoByID(idx, &todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Error": err.Error()})
		return
	}
	fmt.Println("todo:", todo)
	c.BindJSON(&todo)
	if err := dao.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"Error": err.Error()})
		return
	}
}
