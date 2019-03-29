package models

import (
	u "go-todo-rest-api/utils"
	"log"

	"github.com/jinzhu/gorm"
)

// Task represents a task
type Task struct {
	gorm.Model
	UserID uint   `json:"-" gorm:"not_null;index"`
	Name   string `json:"name"`
}

// Create is used to create a new task
func (task *Task) Create() map[string]interface{} {
	GetDB().Create(task)
	if task.ID <= 0 {
		return u.Message(false, "Failed to create task, connection error.")
	}

	response := u.Message(true, "Task has been created")
	response["task"] = task
	return response
}

// Update is used to update a task
func (task *Task) Update() map[string]interface{} {
	GetDB().Save(task)
	if task.ID <= 0 {
		return u.Message(false, "Failed to update task, connection error.")
	}

	response := u.Message(true, "Task has been updated")
	response["task"] = task
	return response
}

// Delete is used to delete the provided task
func (task *Task) Delete() map[string]interface{} {
	GetDB().Delete(task)
	if task.ID <= 0 {
		return u.Message(false, "Failed to delete task, connection error.")
	}

	response := u.Message(true, "Task has been deleted")
	return response
}

// GetTaskByID will return task by provided id
func GetTaskByID(id uint) *Task {
	task := &Task{}
	GetDB().Table("tasks").Where("id = ?", id).First(task)
	if task.Name == "" { //User not found!
		return nil
	}

	return task
}

// GetTasksByUserID will return tasks by provided userID
func GetTasksByUserID(userID uint) []*Task {
	tasks := make([]*Task, 0)
	err := GetDB().Table("tasks").Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		log.Panic(err)
	}
	return tasks
}
