package project

import (
	"encoding/json"
	"net/http"

	"golang-mux-gorm-boilerplate/app/service"

	"golang-mux-gorm-boilerplate/app/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

const MODEL = "Project"

func GetAllProjects(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	projects := []model.Project{}
	db.Find(&projects)
	service.RespondJSON(w, http.StatusOK, projects)
}

func GetProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := service.GetByFieldOr404(db, w, r, &model.Project{Title: title}, MODEL)
	if project == nil {
		return
	}
	service.RespondJSON(w, http.StatusOK, project)
}

func CreateProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	project := model.Project{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		service.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := service.CreateUpdateOr404(db, w, r, project); err != nil {
		service.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	service.RespondJSON(w, http.StatusCreated, project)
}

func UpdateProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := service.GetByFieldOr404(db, w, r, &model.Project{Title: title}, MODEL)
	if project == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		service.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := service.CreateUpdateOr404(db, w, r, *project); err != nil {
		service.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	service.RespondJSON(w, http.StatusOK, project)
}

func DeleteProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := service.GetByFieldOr404(db, w, r, &model.Project{Title: title}, MODEL)
	if project == nil {
		return
	}
	service.DeleteOr404(db, w, r, *project)
	service.RespondJSON(w, http.StatusNoContent, nil)
}

func ArchiveProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := service.GetByFieldOr404(db, w, r, &model.Project{Title: title}, MODEL)
	if project == nil {
		return
	}
	// project.Archive()
	if err := service.CreateUpdateOr404(db, w, r, *project); err != nil {
		service.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	service.RespondJSON(w, http.StatusOK, project)
}

func RestoreProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := service.GetByFieldOr404(db, w, r, &model.Project{Title: title}, MODEL)
	if project == nil {
		return
	}
	// project.Restore()
	if err := service.CreateUpdateOr404(db, w, r, *project); err != nil {
		service.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	service.RespondJSON(w, http.StatusOK, project)
}
