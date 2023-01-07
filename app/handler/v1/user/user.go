package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang-mux-gorm-boilerplate/app/handler"

	"golang-mux-gorm-boilerplate/app/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []model.User{}
	db.Find(&users)
	handler.RespondJSON(w, http.StatusOK, users)
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&user).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusCreated, user)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, _ := strconv.Atoi(vars["id"])
	user := getUserOr404(db, userId, w, r)
	if user == nil {
		return
	}

	handler.RespondJSON(w, http.StatusOK, user)
}

func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, _ := strconv.Atoi(vars["id"])
	user := getUserOr404(db, userId, w, r)
	if user == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&user).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusOK, user)
}

func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, _ := strconv.Atoi(vars["id"])
	user := getUserOr404(db, userId, w, r)
	if user == nil {
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusNoContent, nil)
}

// getUserOr404 gets a user instance if exists, or respond the 404 error otherwise
func getUserOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *model.User {
	user := model.User{}
	if err := db.First(&user, id).Error; err != nil {
		handler.RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}
