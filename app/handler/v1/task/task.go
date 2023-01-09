package task

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang-mux-gorm-boilerplate/app/service"

	"golang-mux-gorm-boilerplate/app/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

const MODEL = "Task"

func GetAllTasks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := service.GetByFieldOr404(db, w, r, &model.Project{Title: projectTitle}, "Project")
	if project == nil {
		return
	}

	tasks := []model.Task{}
	if err := db.Model(&project).Related(&tasks).Error; err != nil {
		service.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	service.RespondJSON(w, http.StatusOK, tasks)
}

func GetTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := service.GetByFieldOr404(db, w, r, &model.Project{Title: projectTitle}, "Project")
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	task := service.GetByIdOr404(db, w, r, id, MODEL)
	if task == nil {
		return
	}
	service.RespondJSON(w, http.StatusOK, task)
}

func CreateTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := service.GetByFieldOr404(db, w, r, &model.Project{Title: projectTitle}, "Project")
	if project == nil {
		return
	}

	// task := model.Task{ProjectID: project}
	task := model.Task{ProjectID: 1}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		service.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := service.CreateUpdateOr404(db, w, r, project); err != nil {
		service.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	service.RespondJSON(w, http.StatusCreated, task)
}

func UpdateTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := service.GetByFieldOr404(db, w, r, &model.Project{Title: projectTitle}, "Project")
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	task := service.GetByIdOr404(db, w, r, id, MODEL)
	if task == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		service.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := service.CreateUpdateOr404(db, w, r, project); err != nil {
		service.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	service.RespondJSON(w, http.StatusOK, task)
}

func DeleteTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := service.GetByFieldOr404(db, w, r, &model.Project{Title: projectTitle}, "Project")
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	task := service.GetByIdOr404(db, w, r, id, MODEL)
	if task == nil {
		return
	}

	service.DeleteOr404(db, w, r, *task)
	service.RespondJSON(w, http.StatusNoContent, nil)
}

func CompleteTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := service.GetByFieldOr404(db, w, r, &model.Project{Title: projectTitle}, "Project")
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	task := service.GetByIdOr404(db, w, r, id, MODEL)
	if task == nil {
		return
	}

	// task.Complete()
	if err := service.CreateUpdateOr404(db, w, r, project); err != nil {
		service.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	service.RespondJSON(w, http.StatusOK, task)
}

func UndoTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := service.GetByFieldOr404(db, w, r, &model.Project{Title: projectTitle}, "Project")
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	task := service.GetByIdOr404(db, w, r, id, MODEL)
	if task == nil {
		return
	}

	// task.Undo()
	if err := service.CreateUpdateOr404(db, w, r, project); err != nil {
		service.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	service.RespondJSON(w, http.StatusOK, task)
}
