package rcache

import (
	"encoding/json"
	"errors"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	redisPool     *redis.Pool
	redisServer   = "127.0.0.1:6379"
	redisPassword = "2a8RQ96jSwCIp9jWEzyPDUy8CSDx3db5FMNh6RSu2FXi3pIOGR5kox11SnmNTbOxkPXqUzOA8ytaH61Q"
)
var ErrCantConnect = errors.New("Can't connect to redis")

func dial() (redis.Conn, error) {
	if redisPool == nil {
		redisPool = &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", redisServer)
				if err != nil {
					return nil, err
				}
				if _, err := c.Do("AUTH", redisPassword); err != nil {
					c.Close()
					return nil, err
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		}
	}
	if redisPool != nil {
		conn := redisPool.Get()
		return conn, nil
	}
	return nil, ErrCantConnect
}

func Set(key string, value interface{}) error {
	c, err := dial()
	if err != nil {
		return err
	}
	defer c.Close()

	jsonvalue, err := json.MarshalIndent(value, " ", " ")
	if err != nil {
		return err
	}

	c.Do("SET", key, jsonvalue, "EX", 120)
	return nil
}

func Get(key string, entityPointer interface{}) error {
	c, err := dial()
	if err != nil {
		return err
	}
	defer c.Close()

	jsonvalue, err := redis.String(c.Do("GET", key))
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(jsonvalue), entityPointer)
	return err
}
