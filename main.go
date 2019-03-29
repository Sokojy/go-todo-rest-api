package main

import (
	"fmt"
	"go-todo-rest-api/controllers"
	"go-todo-rest-api/middleware"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(middleware.JwtAuthentication) //attach JWT auth middleware

	router.HandleFunc("/api/user/new", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/task/new", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/api/task/{id:[0-9]+}/update", controllers.UpdateTask).Methods("POST")
	router.HandleFunc("/api/task/{id:[0-9]+}/delete", controllers.DeleteTask).Methods("POST")
	router.HandleFunc("/api/tasks", controllers.GetTasks).Methods("GET")

	port := os.Getenv("PORT") //Get port from .env file
	if port == "" {
		port = "8000"
	}

	fmt.Println("Listen and serve on", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
