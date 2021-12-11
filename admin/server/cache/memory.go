package cache

import (
	"admin/config"
	"bytes"
	"encoding/gob"
	"fmt"
	gocache "github.com/patrickmn/go-cache"
	"time"
)

type Memory struct {
	config  *config.Memory
	memeory *gocache.Cache
}

func setupMemory(cfg *config.Memory) (*Memory, error) {
	memoryCache := new(Memory)
	memoryCache.config = cfg
	memoryCache.memeory = gocache.New(5*time.Minute, memoryCache.config.PurgeTime)

	return memoryCache, nil
}

func (m *Memory) Get(key string) (interface{}, error) {
	if x, found := m.memeory.Get(key); found {
		return x, nil
	}
	return nil, nil
}

func (m *Memory) Set(key string, value interface{}, expire time.Duration) error {
	if expire < 0 {
		return fmt.Errorf("%s", "invalid ttl value")
	}

	buf := bytes.Buffer{}
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(value)
	if err != nil {
		return err
	}

	m.memeory.Set(key, buf.Bytes(), expire)
	return nil
}
