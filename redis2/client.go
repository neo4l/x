package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	Default_MaxIdle   = 5
	Default_MaxActive = 40
	Default_Timeout   = 180 * time.Second
)

type RedisClient struct {
	redisPool *redis.Pool
}

func NewClient(host, passwd string, db int) *RedisClient {
	return NewClient2(host, passwd, db, Default_MaxIdle, Default_MaxActive, Default_Timeout)
}

func NewClient2(host, passwd string, db, maxIdle, maxActive int, timeout time.Duration) *RedisClient {
	pool := &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: timeout,
		Dial: func() (redis.Conn, error) {
			options := []redis.DialOption{redis.DialPassword(passwd), redis.DialDatabase(db)}
			c, err := redis.Dial("tcp", host, options...)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	return &RedisClient{pool}
}

func (self *RedisClient) Get(key string) (string, error) {
	rc := self.redisPool.Get()
	defer rc.Close()

	return redis.String(rc.Do("GET", key))
}

func (self *RedisClient) Exists(key string) (bool, error) {
	rc := self.redisPool.Get()
	defer rc.Close()

	return redis.Bool(rc.Do("EXISTS", key))
}

func (self *RedisClient) Set(key string, value string) (int, error) {
	rc := self.redisPool.Get()
	defer rc.Close()

	return redis.Int(rc.Do("SET", key, value))
}

func (self *RedisClient) SetWithExpire(key string, value string, expireTime int) (int, error) {
	rc := self.redisPool.Get()
	defer rc.Close()

	return redis.Int(rc.Do("SET", key, value, "EX", expireTime))
}
