package routes

import (
	"encoding/json"
	"net/http"

	"github.com/JuanM-Ortiz/go-gorm-api-rest/db"
	"github.com/JuanM-Ortiz/go-gorm-api-rest/models"
	"github.com/gorilla/mux"
)

func GetAllTasks(rw http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)

	json.NewEncoder(rw).Encode(&tasks)
}

func GetTask(rw http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Task Not Found"))
		return
	}

	json.NewEncoder(rw).Encode(&task)

}

func CreateTask(rw http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	createdTask := db.DB.Create(&task)
	err := createdTask.Error

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(rw).Encode(&task)
}

func UpdateTask(rw http.ResponseWriter, r *http.Request) {

}

func DeleteTask(rw http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Task Not Found"))
		return
	}

	db.DB.Unscoped().Delete(&task)
	rw.WriteHeader(http.StatusNoContent) // 204
}
