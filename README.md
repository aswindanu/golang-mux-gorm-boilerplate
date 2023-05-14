# Go Todo REST API Example
A RESTful API example for simple todo application with Go

It is a just simple tutorial or example for making simple RESTful API with Go using **gorilla/mux** (A nice mux library) and **gorm** (An ORM for Go)

## Installation
```bash
# Download this project
go get github.com/aswindanu/golang-mux-gorm-boilerplate
```

Before running API server, you should set the database config with yours or set the your database config with my values on [config.go](https://golang-mux-gorm-boilerplate/blob/master/config/config.go)

```bash
1. <Make a Postgresql database connection (server / localhost)>
2. cp .env .env.sample
3. <fill necessary env field>
```

for Development, we are using hot reload by installing nodemon.
see docs https://techinscribed.com/5-ways-to-live-reloading-go-applications/

##Run
```bash
DEVELOPMENT # API Endpoint : http://127.0.0.1:3003
- docker-compose up -d --build backend-dev


PRODUCTION # API Endpoint : http://127.0.0.1:3000
# Build and Run
- docker-compose up -d --build backend
```

## Structure
```
├── app
│   ├── app.go
│   ├── handler          	// Our API core handlers
│   │   ├── common.go    	// Common response functions
│   │   ├── projects
│   │   │   └── projects.go // APIs for Project model
│   │   ├── tasks
│   │   │   └── tasks.go    // APIs for Task model
│   │   └── ... (and so on)
│   │
│   └── model
│       ├── model.go     	// Main models for our application
│       ├── project.go     	// Project model go
│       ├── task.go     	// Task model go
│       └── ... (and so on)
│
├── config
│   └── config.go        	// Configuration (this is mostly related with .env)
│
└── main.go
```

## API

#### /api/v1/go/projects
* `GET` : Get all projects
* `POST` : Create a new project

#### /api/v1/go/projects/:title
* `GET` : Get a project
* `PUT` : Update a project
* `DELETE` : Delete a project

#### /api/v1/go/projects/:title/archive
* `PUT` : Archive a project
* `DELETE` : Restore a project 

#### /api/v1/go/projects/:title/tasks
* `GET` : Get all tasks of a project
* `POST` : Create a new task in a project

#### /api/v1/go/projects/:title/tasks/:id
* `GET` : Get a task of a project
* `PUT` : Update a task of a project
* `DELETE` : Delete a task of a project

#### /api/v1/go/projects/:title/tasks/:id/complete
* `PUT` : Complete a task of a project
* `DELETE` : Undo a task of a project

## Todo

- [x] Support basic REST APIs.
- [ ] Support Authentication with user for securing the APIs.
- [ ] Make convenient wrappers for creating API handlers.
- [ ] Write the tests for all APIs.
- [x] Organize the code with packages
- [ ] Make docs with GoDoc
- [ ] Building a deployment process 
