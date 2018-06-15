package redis

import (
	"andui/conf"
	//"fmt"
	"time"

	"github.com/go-redis/redis"
)

func NewClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Redis_Host,
		Password: conf.Redis_Password, // no password set
		DB:       0,                   // use default DB
	})

	_, err := client.Ping().Result()
	//fmt.Println(pong, err)

	return client, err
}

func Get(key string) (string, error) {
	rc, err := NewClient()
	//fmt.Printf("newClient: %s\n", err)
	if err != nil {
		return "", err
	}
	defer rc.Close()

	return rc.Get(key).Result()
}

func Set(key string, value interface{}, expir time.Duration) error {
	rc, err := NewClient()
	if err != nil {
		return err
	}

	defer rc.Close()

	return rc.Set(key, value, expir).Err()
}

func Expire(key string, expir time.Duration) error {
	rc, err := NewClient()
	if err != nil {
		return err
	}
	defer rc.Close()

	return rc.Expire(key, expir).Err()
}
