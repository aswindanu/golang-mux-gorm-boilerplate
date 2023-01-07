package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Project struct {
	gorm.Model
	Title    string `gorm:"unique" json:"title"`
	Archived bool   `json:"archived"`
	Tasks    []Task `gorm:"ForeignKey:ProjectID" json:"tasks"`
}

func (p *Project) Archive() {
	p.Archived = true
}

func (p *Project) Restore() {
	p.Archived = false
}
