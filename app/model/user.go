package model

import (
	"fmt"
	"net/http"
	"strings"
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
	IpAddress string     `mapstructure:"ip_address" json:"ip_address"`
	CreatedAt time.Time  `mapstructure:"created_at" json:"created_at"`
	UpdatedAt time.Time  `mapstructure:"updated_at" json:"updated_at"`
}

func (u *User) Male() {
	u.Gender = MALE
}

func (u *User) Female() {
	u.Gender = FEMALE
}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func (u *User) ValidateData(w http.ResponseWriter, r *http.Request) error {
	gender := strings.ToUpper(fmt.Sprintf("%v", u.Gender))
	switch gender {
	case "M":
		u.Male()
	case "F":
		u.Female()
	}
	u.Active = true
	u.IpAddress = ReadUserIP(r)
	return nil
}

func (u *User) TableName() string {
	return "user"
}
