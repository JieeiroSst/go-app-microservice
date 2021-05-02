package bigcache

import (
	"github.com/allegro/bigcache"
	"net/http"
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

func (c *Cache) Get(cookie *http.Cookie)([]byte,error){
	bytes,err:=c.cache.Get(cookie.Name)
	return bytes,err
}

func (c *Cache) Iterator() *bigcache.EntryInfoIterator{
	return c.cache.Iterator()
}

func (c *Cache) Delete(key string) error{
	return c.cache.Delete(key)
}

func (c *Cache) Set(key string,bytes []byte)error{
	return c.cache.Set(key,bytes)
}