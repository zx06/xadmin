package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account        string `json:"account,omitempty"`
	PasswordDigest string `json:"password_digest,omitempty"`
	Nickname       string `json:"nickname,omitempty"`
	Status         string `json:"status,omitempty"`
}
