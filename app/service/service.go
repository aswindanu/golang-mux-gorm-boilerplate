package service

import (
	"net/http"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"golang-mux-gorm-boilerplate/app/model"

	"github.com/jinzhu/gorm"
)

// THIS IS BUSINESS CORE QUERY, PLEASE CODE WISELY!
// We're implementing hexagonal architecture to make sure business core keep it's
// best to handle most of the queries

// getAll gets all instances
func GetAll(db *gorm.DB, w http.ResponseWriter, r *http.Request, modelName string) *interface{} {
	data := model.Models()["list"][cases.Title(language.English).String(modelName)]
	if err := db.Find(data).Error; err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &data
}

// GetByIdOr404 gets instance by id if exists, or respond the 404 error otherwise
func GetByIdOr404(db *gorm.DB, w http.ResponseWriter, r *http.Request, id int, modelName string) *interface{} {
	data := model.Models()["get"][cases.Title(language.English).String(modelName)]
	if err := db.First(data, id).Error; err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &data
}

// GetByFieldOr404 gets instance by id if exists, or respond the 404 error otherwise
func GetByFieldOr404(db *gorm.DB, w http.ResponseWriter, r *http.Request, field interface{}, modelName string) *interface{} {
	data := model.Models()["get"][cases.Title(language.English).String(modelName)]
	if err := db.First(data, field).Error; err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &data
}

// CreateUpdateOr404 delete instance if exists, or respond the 404 error otherwise
func CreateUpdateOr404(db *gorm.DB, w http.ResponseWriter, r *http.Request, data interface{}) error {
	if err := db.Save(data).Error; err != nil {
		return err
	}
	return nil
}

// DeleteOr404 delete instance if exists, or respond the 404 error otherwise
func DeleteOr404(db *gorm.DB, w http.ResponseWriter, r *http.Request, data interface{}) error {
	if err := db.Delete(data).Error; err != nil {
		return err
	}
	return nil
}
