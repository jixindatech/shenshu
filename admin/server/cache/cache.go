package cache

import (
	"admin/config"
	"fmt"
	"time"
)

const (
	CONFIG = "config"
	MEMORY = "memory"
	REDIS  = "redis"
)

type Cache interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}, ttl time.Duration) error
}

var cacheItems map[string]interface{}
var cacheConfig string

func SetupCache(cfg *config.Config) error {
	cacheItems = make(map[string]interface{})

	cacheRedis, err := setupRedis(cfg.Redis)
	if err != nil {
		return err
	}
	cacheItems[REDIS] = cacheRedis

	memory, err := setupMemory(cfg.Memory)
	if err != nil {
		return err
	}
	cacheItems[MEMORY] = memory

	if cfg.Cache == REDIS {
		cacheConfig = REDIS
	} else {
		return fmt.Errorf("%s", "invalid cache type")
	}

	return nil
}

func Get(cacheType string, key string) (interface{}, error) {
	if cacheType == CONFIG {
		cacheType = cacheConfig
	}
	instance, ok := cacheItems[cacheType]
	if ok {
		return instance.(Cache).Get(key)
	}

	return nil, fmt.Errorf("%s", "unknown cache type")
}

// 0 forever
func Set(cacheType string, key string, value interface{}, ttl time.Duration) error {
	if ttl < 0 {
		return fmt.Errorf("%s", "invalid ttl value")
	}

	if cacheType == CONFIG {
		cacheType = cacheConfig
	}

	instance, ok := cacheItems[cacheType]
	if ok {
		return instance.(Cache).Set(key, value, ttl)
	}

	return fmt.Errorf("%s", "unknown cache type")
}
