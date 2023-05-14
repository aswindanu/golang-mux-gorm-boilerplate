package beratBadan

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang-mux-gorm-boilerplate/app/handler"

	"golang-mux-gorm-boilerplate/app/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllBeratBadan(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	beratBadan := []model.BeratBadan{}
	db.Find(&beratBadan)
	handler.RespondJSON(w, http.StatusOK, beratBadan)
}

func CreateBeratBadan(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	beratBadan := model.BeratBadan{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&beratBadan); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&beratBadan).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusCreated, beratBadan)
}

func GetBeratBadan(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	beratBadanId, _ := strconv.Atoi(vars["id"])
	beratBadan := getBeratBadanOr404(db, beratBadanId, w, r)
	if beratBadan == nil {
		return
	}

	handler.RespondJSON(w, http.StatusOK, beratBadan)
}

func UpdateBeratBadan(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	beratBadanId, _ := strconv.Atoi(vars["id"])
	beratBadan := getBeratBadanOr404(db, beratBadanId, w, r)
	if beratBadan == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&beratBadan); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&beratBadan).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusOK, beratBadan)
}

func DeleteBeratBadan(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	beratBadanId, _ := strconv.Atoi(vars["id"])
	beratBadan := getBeratBadanOr404(db, beratBadanId, w, r)
	if beratBadan == nil {
		return
	}

	if err := db.Delete(&beratBadan).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusNoContent, nil)
}

// getBeratBadanOr404 gets a beratBadan instance if exists, or respond the 404 error otherwise
func getBeratBadanOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *model.BeratBadan {
	beratBadan := model.BeratBadan{}
	if err := db.First(&beratBadan, id).Error; err != nil {
		handler.RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &beratBadan
}
