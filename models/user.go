package models

import (
	"time"
)

type User struct {
	IdUser    uint   `gorm:"primary_key" json:"iduser"`
	NamaUser  string `json:"namauser"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	IsActive  int    `json:"isactive"`
	NoCif     string `json:"nocif"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Addresses []Address `json:"addresses" gorm:"foreignkey:IdUser"`
}

type Address struct {
	IdAddress    uint   `gorm:"primary_key" json:"idalamat"`
	IdUser       uint   `json:"iduser"`
	AlamatDetail string `json:"alamatdetail"`
	Provinsi     string `json:"provinsi"`
	Kabupaten    string `json:"kabupaten"`
	KodePos      string `json:"kodepos"`
}

type Log struct {
	IdLog       uint      `gorm:"primary_key" json:"idlog"`
	IdUser      uint      `json:"iduser"`
	IsLogin     int       `json:"islogin"`
	ExpiredLock time.Time `json:"expiredlock"`
	Email       string    `json:"email"`
}
