package config

import (
	"fmt"
	"sync"

	"github.com/caarlos0/env/v6"
)

var (
	cfg     = &conf{}
	cfgOnce sync.Once
)

type conf struct {
	Port        int    `env:"PORT" envDefault:"3000"`
	DatabaseURL string `env:"DATABASE_URL"`
	RedisURL    string `env:"REDIS_URL"`
}

// Config 单例,所有设置都从这里取
func Config() *conf {
	cfgOnce.Do(func() {
		if err := env.Parse(cfg); err != nil {
			fmt.Printf("%+v\n", err)
		}
	})
	return cfg
}
