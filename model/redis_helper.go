package model

import (
	// "fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/op/go-logging"
	"time"
)

var (
	log       = logging.MustGetLogger("model")
	redisPool *redis.Pool
	// redisServer string = "10.10.79.123:15151" // host:port of server
	redisServer string = "localhost:6379" // host:port of server
	password    string
)

func init() {
	redisPool = &redis.Pool{
		MaxIdle:     3,
		MaxActive:   160,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisServer)
			if err != nil {
				return nil, err
			}
			// if _, err := c.Do("AUTH", password); err != nil {
			// 	c.Close()
			// 	return nil, err
			// }
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
