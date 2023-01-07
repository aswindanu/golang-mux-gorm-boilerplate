package model

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type BeratBadan struct {
	Id        int       `gorm:"primary_key"; "AUTO_INCREMENT" mapstructure:"id" json:"id"`
	Max       int       `json:"max"`
	Min       int       `json:"min"`
	Perbedaan int       `json:"perbedaan"`
	Tanggal   time.Time `json:"tanggal"`
}

func (b *BeratBadan) TableName() string {
	return "berat_badan"
}
