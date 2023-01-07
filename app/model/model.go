package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBModels struct {
	User       interface{}
	BeratBadan interface{}
	Project    interface{}
	Task       interface{}
}

func Models() *DBModels {
	return &DBModels{
		&User{},
		&BeratBadan{},
		&Project{},
		&Task{},
	}
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(
		&User{},
		&BeratBadan{},
		&Project{},
		&Task{},
	)
	// db.Model(&Task{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	return db
}
