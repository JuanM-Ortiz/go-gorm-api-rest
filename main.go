package main

import (
	"net/http"

	"github.com/JuanM-Ortiz/go-gorm-api-rest/db"
	"github.com/JuanM-Ortiz/go-gorm-api-rest/models"
	"github.com/JuanM-Ortiz/go-gorm-api-rest/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DbConnect()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	//Users routes

	router.HandleFunc("/users", routes.GetAllUsers).Methods("GET")
	router.HandleFunc("/users", routes.PostUser).Methods("POST")
	router.HandleFunc("/users/{id}", routes.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", routes.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", routes.DeleteUser).Methods("DELETE")

	//Tasks routes

	router.HandleFunc("/tasks", routes.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks", routes.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", routes.DeleteTask).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
