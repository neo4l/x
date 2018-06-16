package redis

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	client *redis.Client
}

func NewClient(host, passwd string) (*RedisClient, error) {
	return NewClient2(host, passwd, 0)
}

func NewClient2(host, passwd string, db int) (*RedisClient, error) {
	newClient := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: passwd, // no password set
		DB:       db,     // use default DB
	})

	_, err := newClient.Ping().Result()
	//fmt.Println(pong, err)
	if err != nil {
		return nil, err
	}
	return &RedisClient{client: newClient}, err
}

func (self *RedisClient) Get(key string) (string, error) {
	return self.client.Get(key).Result()
}

func (self *RedisClient) Set(key string, value interface{}, expir time.Duration) error {
	return self.client.Set(key, value, expir).Err()
}

func (self *RedisClient) Expire(key string, expir time.Duration) error {
	return self.client.Expire(key, expir).Err()
}

func (self *RedisClient) Close() error {
	return self.client.Close()
}
