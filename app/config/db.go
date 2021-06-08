package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func GORMFactory(dsn string, gormConf *gorm.Config) *gorm.DB {
	db, err := gorm.Open(
		postgres.Open(dsn),
		gormConf,
	)
	if err != nil {
		log.Panicln(err)
	}
	db.Exec("select 1")
	return db
}

func DB() *gorm.DB {
	dbOnce.Do(func() {
		db = GORMFactory(cfg.DatabaseURL, &gorm.Config{})
	})
	return db
}
