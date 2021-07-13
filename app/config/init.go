package config

import (
	"log"
	"xadmin/app/model"
)

func init() {
	// 自动迁移
	err := DB().AutoMigrate(&model.User{})
	if err != nil {
		log.Panicln(err)
	}
}
