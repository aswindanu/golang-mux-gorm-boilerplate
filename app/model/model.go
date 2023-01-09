package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Models() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"list": {
			"User":       &[]User{},
			"BeratBadan": &[]BeratBadan{},
			"Project":    &[]Project{},
			"Task":       &[]Task{},
		},
		"get": {
			"User":       &User{},
			"BeratBadan": &BeratBadan{},
			"Project":    &Project{},
			"Task":       &Task{},
		},
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
