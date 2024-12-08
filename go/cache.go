package main

import (
	"errors"
	"sync"
	"time"
)

var ErrCacheNotFound = errors.New("cache not found")
var ErrCacheExpired = errors.New("cache expired")

type Cache interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, error)
}

type CacheInMemoryCfg struct {
	cacheClearTime time.Duration
}

type CacheInMemory struct {
	store map[string]*valueWithTime
	sync  *sync.Mutex
	*CacheInMemoryCfg
}

type valueWithTime struct {
	value interface{}
	time  time.Time
}

var defaultCacheInMemoryCfg = &CacheInMemoryCfg{
	cacheClearTime: time.Hour,
}

func NewCacheInMemory(config *CacheInMemoryCfg) *CacheInMemory {
	cache := &CacheInMemory{
		store: make(map[string]*valueWithTime),
		sync:  &sync.Mutex{},
	}
	if config == nil {
		config = defaultCacheInMemoryCfg
	}
	cache.CacheInMemoryCfg = config
	return cache
}

func (c *CacheInMemory) Set(key string, value interface{}) {
	c.sync.Lock()
	defer c.sync.Unlock()
	c.store[key] = &valueWithTime{
		value: value,
		time:  time.Now(),
	}
}

func (c *CacheInMemory) Get(key string) (interface{}, error) {
	c.sync.Lock()
	defer c.sync.Unlock()

	v, ok := c.store[key]
	if !ok {
		return nil, ErrCacheNotFound
	}

	diff := time.Since(v.time)
	if diff > c.cacheClearTime {
		delete(c.store, key)
		return nil, ErrCacheExpired
	}
	return v.value, nil
}

func (c *CacheInMemory) Clear(key string) {
	c.sync.Lock()
	defer c.sync.Unlock()
	delete(c.store, key)
}


var userCache = NewCacheInMemory(nil)
