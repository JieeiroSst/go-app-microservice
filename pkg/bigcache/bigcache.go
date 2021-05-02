package bigcache

import (
	"github.com/allegro/bigcache"
	"time"
)

type Cache struct {
	cache *bigcache.BigCache
}


func NewCache(cache *bigcache.BigCache) *Cache{
	return &Cache{cache:cache}
}

func init() {
	_,_=bigcache.NewBigCache(bigcache.DefaultConfig(30*time.Hour))
}

func (c *Cache) Get(key string)([]byte,error){
	bytes,err:=c.cache.Get(key)
	return bytes,err
}