package app

import (
	"fmt"
	"log"
	"net/http"

	// module
	"golang-mux-gorm-boilerplate/app/handler/v1/project"
	"golang-mux-gorm-boilerplate/app/handler/v1/task"
	"golang-mux-gorm-boilerplate/app/handler/v1/user"

	// model
	"golang-mux-gorm-boilerplate/app/model"

	// config
	"golang-mux-gorm-boilerplate/config"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	var err error
	var db *gorm.DB
	var dbURI string

	// MYSQL
	if config.DB.Dialect == "mysql" {
		dbURI = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
			config.DB.Username,
			config.DB.Password,
			config.DB.Host,
			config.DB.Port,
			config.DB.Name,
			config.DB.Charset,
		)
	}
	// POSTGRES
	if config.DB.Dialect == "postgres" {
		dbURI = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			config.DB.Host,
			config.DB.Username,
			config.DB.Password,
			config.DB.Name,
			config.DB.Port,
		)
	}

	db, err = gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		fmt.Println("dbURI: ", dbURI)
		log.Fatal("Could not connect database:" + err.Error())
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	// Router v1
	a.routerV1()
}

func (a *App) endpointAPI(uri string, endpoint string) string {
	return fmt.Sprintf("%v/%v", uri, endpoint)
}

func (a *App) routerV1() {
	uri := "/api/v1/go"

	// User
	a.Get(a.endpointAPI(uri, "user"), a.handleRequest(user.GetAllUsers))
	a.Post(a.endpointAPI(uri, "user"), a.handleRequest(user.CreateUser))
	a.Get(a.endpointAPI(uri, "user/{id}"), a.handleRequest(user.GetUser))
	a.Put(a.endpointAPI(uri, "user/{id}"), a.handleRequest(user.UpdateUser))
	a.Delete(a.endpointAPI(uri, "user/{id}"), a.handleRequest(user.DeleteUser))

	// Project
	a.Get(a.endpointAPI(uri, "projects"), a.handleRequest(project.GetAllProjects))
	a.Post(a.endpointAPI(uri, "projects"), a.handleRequest(project.CreateProject))
	a.Get(a.endpointAPI(uri, "projects/{title}"), a.handleRequest(project.GetProject))
	a.Put(a.endpointAPI(uri, "projects/{title}"), a.handleRequest(project.UpdateProject))
	a.Delete(a.endpointAPI(uri, "projects/{title}"), a.handleRequest(project.DeleteProject))
	a.Put(a.endpointAPI(uri, "projects/{title}/archive"), a.handleRequest(project.ArchiveProject))
	a.Delete(a.endpointAPI(uri, "projects/{title}/archive"), a.handleRequest(project.RestoreProject))

	// Tasks
	a.Get(a.endpointAPI(uri, "projects/{title}/tasks"), a.handleRequest(task.GetAllTasks))
	a.Post(a.endpointAPI(uri, "projects/{title}/tasks"), a.handleRequest(task.CreateTask))
	a.Get(a.endpointAPI(uri, "projects/{title}/tasks/{id:[0-9]+}"), a.handleRequest(task.GetTask))
	a.Put(a.endpointAPI(uri, "projects/{title}/tasks/{id:[0-9]+}"), a.handleRequest(task.UpdateTask))
	a.Delete(a.endpointAPI(uri, "projects/{title}/tasks/{id:[0-9]+}"), a.handleRequest(task.DeleteTask))
	a.Put(a.endpointAPI(uri, "projects/{title}/tasks/{id:[0-9]+}/complete"), a.handleRequest(task.CompleteTask))
	a.Delete(a.endpointAPI(uri, "projects/{title}/tasks/{id:[0-9]+}/complete"), a.handleRequest(task.UndoTask))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}
