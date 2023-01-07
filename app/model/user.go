package model

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type genderType string

const (
	MALE   genderType = "M"
	FEMALE genderType = "F"
)

type User struct {
	Id        int        `gorm:"primary_key"; "AUTO_INCREMENT" mapstructure:"id" json:"id"`
	Email     string     `gorm:"unique" json:"email"`
	Username  string     `gorm:"unique" json:"username"`
	Password  string     `json:"password"`
	Fullname  string     `json:"fullname"`
	Phone     string     `json:"phone"`
	Gender    genderType `gorm:"type:gender_type" json:"gender"`
	Active    bool       `gorm:"default:true" json:"active"`
	IpAddress string     `gorm:"default:true" mapstructure:"ip_address" json:"ip_address"`
	CreatedAt time.Time  `mapstructure:"created_at" json:"created_at"`
	UpdatedAt time.Time  `mapstructure:"updated_at" json:"updated_at"`
}

func (u *User) Male() {
	u.Gender = MALE
}

func (u *User) Female() {
	u.Gender = FEMALE
}

func (u *User) TableName() string {
	return "user"
}
