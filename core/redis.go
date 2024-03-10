package core

import (
	"xgo/config"

	"git.qdreads.com/gotools/redis"
)

func init() {
	initRedis()
}

var redisMap = map[string]*redis.RedisClient{}

// redis 实例类型
const (
	RDS_DEFAULT = "default"
)

func initRedis() {
	rdsCfg := config.Cfg.Redis
	for key, cfg := range rdsCfg {
		redisMap[key] = redis.InitRedis(redis.Options{
			Host:     cfg.Host,
			DB:       cfg.DB,
			Password: cfg.Password,
			Mode:     cfg.Mode,
		})
	}
}

func Redis(instance ...string) *redis.RedisClient {
	key := RDS_DEFAULT
	if len(instance) > 0 {
		key = instance[0]
	}

	return redisMap[key]
}
