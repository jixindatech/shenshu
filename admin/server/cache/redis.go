package cache

import (
	"admin/config"
	"admin/core/log"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"time"
)

var redisPool *redis.Pool

type Redis struct {
	config *config.Redis
	pool   *redis.Pool
}

func setupRedis(cfg *config.Redis) (*Redis, error) {
	redisCache := new(Redis)
	redisCache.config = cfg
	redisPool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   10,
		IdleTimeout: 240,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp",
				cfg.Host+fmt.Sprintf(":%d", cfg.Port),
				redis.DialDatabase(cfg.Db),
				redis.DialReadTimeout(time.Duration(1)*time.Second),
				redis.DialWriteTimeout(time.Duration(1)*time.Second),
				redis.DialConnectTimeout(time.Duration(2)*time.Second),
			)
			if err != nil {
				return nil, err
			}
			if cfg.Password != "" {
				if _, err := c.Do("AUTH", cfg.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if _, err := c.Do("PING"); err != nil {
				return err
			}
			return nil
		},
	}

	redisCache.pool = redisPool
	return redisCache, nil
}

func (r *Redis) Set(key string, value interface{}, ttl time.Duration) error {
	if r.config == nil {
		return fmt.Errorf("%s", "invalid redis instance")
	}

	conn := redisPool.Get()
	defer conn.Close()

	if r.config.KeyPrefix != "" {
		key = r.config.KeyPrefix + key
	}

	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(value)
	if err != nil {
		return err
	}

	if ttl == 0 {
		_, err = conn.Do("SET", key, buffer.Bytes())
		if err != nil {
			log.Logger.Error("redis", zap.String("err", err.Error()))
			return err
		}

	} else if ttl > 0 {
		_, err = conn.Do("SET", key, buffer.Bytes(), "EX", int(ttl))
		if err != nil {
			log.Logger.Error("redis", zap.String("err", err.Error()))
			return err
		}
	}

	return err
}

func (r *Redis) Get(key string) (interface{}, error) {
	if r.config == nil {
		return "", fmt.Errorf("%s", "invalid redis instance")
	}

	conn := redisPool.Get()
	defer conn.Close()

	if r.config.KeyPrefix != "" {
		key = r.config.KeyPrefix + key
	}

	return redis.Bytes(conn.Do("GET", key))
}
