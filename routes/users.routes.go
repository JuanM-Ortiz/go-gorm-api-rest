package routes

import (
	"encoding/json"
	"net/http"

	"github.com/JuanM-Ortiz/go-gorm-api-rest/db"
	"github.com/JuanM-Ortiz/go-gorm-api-rest/models"
	"github.com/gorilla/mux"
)

func GetAllUsers(rw http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(rw).Encode(&users)

}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	//fmt.Println(params["id"])

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("User Not Found"))
		return
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(rw).Encode(&user)

}

func PostUser(rw http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest) //400
		rw.Write([]byte(err.Error()))
	}

	json.NewEncoder(rw).Encode(&user)
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("User Not Found"))
		return
	}

	db.DB.Model(&user).Updates(&newUser)

	json.NewEncoder(rw).Encode(&newUser)
	rw.WriteHeader(http.StatusOK)

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("User Not Found"))
		return
	}

	db.DB.Unscoped().Delete(&user)
	rw.WriteHeader(http.StatusOK)
}
