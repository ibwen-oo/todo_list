package models

import "bubble/dao"

type Todo struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func GetAllTodo() (todoList []*Todo, err error) {
	err = dao.DB.Find(&todoList).Error
	return todoList, err
}

func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func ModifyTodoByID(id int64, todo *Todo) (err error) {
	err = dao.DB.Where("id=?", id).First(&todo).Error
	return
}

func DeleteTodoByID(id int64) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return err
}
