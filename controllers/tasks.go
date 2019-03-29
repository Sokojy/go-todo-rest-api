package controllers

import (
	"encoding/json"
	"go-todo-rest-api/models"
	u "go-todo-rest-api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateTask is used to create a new task
//
// /api/task/new
var CreateTask = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	task := &models.Task{}

	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	task.UserID = user
	resp := task.Create()
	u.Respond(w, resp)
}

// UpdateTask is used to update a task
//
// /api/task/{id}/update
var UpdateTask = func(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid task ID"))
		return
	}
	task := models.GetTaskByID(uint(id))
	if userID != task.UserID {
		u.Respond(w, u.Message(false, "Invalid request"))
	}

	decodedTask := &models.Task{}
	err = json.NewDecoder(r.Body).Decode(decodedTask)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	task.Name = decodedTask.Name

	resp := task.Update()
	u.Respond(w, resp)
}

// DeleteTask is used to delete the provided task
//
// /api/task/{id}/delete
var DeleteTask = func(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid task ID"))
		return
	}
	task := models.GetTaskByID(uint(id))
	if userID != task.UserID {
		u.Respond(w, u.Message(false, "Invalid request"))
	}

	resp := task.Delete()
	u.Respond(w, resp)
}

// GetTasks returns all tasks of an authorized user
//
// /api/tasks
var GetTasks = func(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	tasks := []*models.Task{}
	tasks = models.GetTasksByUserID(userID)
	response := u.Message(true, "success")
	response["tasks"] = tasks
	u.Respond(w, response)
}
