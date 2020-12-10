package redis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-session/redis"
	"github.com/go-session/session"
	"github.com/linqiurong2021/gin-arcgis/config"
)

// RedisStore RedisStore
var RedisStore session.Store

// InitRedisSession 初始化
func InitRedisSession(cfg *config.RedisConfig) {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	// 初始化Session
	session.InitManager(
		session.SetStore(redis.NewRedisStore(&redis.Options{
			Addr: addr,
			DB:   cfg.Database,
		})),
	)
}

// GetStore 获取Store
func GetStore(c *gin.Context) (store session.Store, err error) {
	return session.Start(c, c.Writer, c.Request)
}

// Set 设置Redis
func Set(c *gin.Context, key string, value interface{}) (store session.Store, err error) {
	//
	store, err = GetStore(c)
	if err != nil {
		return nil, err
	}
	store.Set(key, value)
	err = store.Save()
	if err != nil {
		return nil, err
	}
	return store, nil
}

// Get 获取数据
func Get(c *gin.Context, key string) (value interface{}, ok bool, err error) {
	//
	store, err := GetStore(c)
	if err != nil {
		return nil, false, err
	}
	value, ok = store.Get(key)
	return value, ok, nil
}
