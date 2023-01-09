package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang-mux-gorm-boilerplate/app/service"

	"golang-mux-gorm-boilerplate/app/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

const MODEL = "User"

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := service.GetAll(db, w, r, MODEL)
	if users == nil {
		return
	}

	service.RespondJSON(w, http.StatusOK, users)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, _ := strconv.Atoi(vars["id"])
	user := service.GetByIdOr404(db, w, r, userId, MODEL)
	if user == nil {
		return
	}

	service.RespondJSON(w, http.StatusOK, user)
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := &model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(user); err != nil {
		service.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	user.ValidateData(w, r)
	if err := service.CreateUpdateOr404(db, w, r, user); err != nil {
		service.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	service.RespondJSON(w, http.StatusCreated, user)
}

func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, _ := strconv.Atoi(vars["id"])
	user := service.GetByIdOr404(db, w, r, userId, MODEL)
	if user == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(user); err != nil {
		service.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := service.CreateUpdateOr404(db, w, r, *user); err != nil {
		service.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	service.RespondJSON(w, http.StatusOK, user)
}

func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, _ := strconv.Atoi(vars["id"])
	user := service.GetByIdOr404(db, w, r, userId, MODEL)
	if user == nil {
		return
	}

	service.DeleteOr404(db, w, r, *user)
	service.RespondJSON(w, http.StatusNoContent, nil)
}
