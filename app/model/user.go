package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account        string `json:"account"`
	Nickname       string `json:"nickname"`
	PasswordDigest string `json:"password_digest"`
	Status         string `json:"status"`
	Admin          bool   `json:"admin" gorm:"default:false"`
}
