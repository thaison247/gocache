package my_gocache

import (
	// "errors"

	"errors"

	"github.com/gomodule/redigo/redis"
)

var (
	// CSConn: connector to redis server
	CSConn *redis.Conn
)

// Connect to redis server
func (r Redis) Connect() error {
	conn, err := redis.Dial("tcp", r.Host+":"+r.Port, redis.DialPassword(r.Password))

	if err != nil {
		return err
	}

	CSConn = &conn
	return nil
}

// Set method
func (r Redis) Set(key string, value interface{}, expireTime ...int) error {
	// set key - value
	_, err := (*CSConn).Do("SET", key, value)

	if err != nil {
		return err
	}

	// set expire time on key
	if len(expireTime) > 0 {
		_, err = (*CSConn).Do("EXPIRE", key, expireTime[0])
		if err != nil {
			return err
		}

		_, err = (*CSConn).Do("TTL", key)
		if err != nil {
			return err
		}
	}

	return nil
}

// Get method
func (r Redis) Get(key string) (interface{}, error) {
	dat, err := redis.String((*CSConn).Do("GET", key))

	if err != nil {
		return nil, err
	}

	return dat, nil
}

// Delete method
func (r Redis) Delete(key string) (int64, error) {
	result, err := (*CSConn).Do("DEL", key)

	return result.(int64), err
}

// Close connection
func (r Redis) Close() {
	(*CSConn).Close()
}

// set expire time on a key
func (r Redis) Expire(key string, expireTime int) (int64, error) {
	val, err := (*CSConn).Do("EXPIRE", key, expireTime)

	if err != nil {
		return 0, err
	}

	return val.(int64), nil
}

// func (r Redis) ExpireV2(key string, expireTime int) error {
// 	val, err := (*CSConn).Do("EXPIRE", key, expireTime)

// 	if val == 0 {
// 		return errors.New("Invalid key")
// 	}

// 	return err
// }

// get remain time life (expiretime)
func (r Redis) GetRemainLifeTime(key string) (int64, error) {
	val, err := (*CSConn).Do("TTL", key)

	if err != nil {
		return 0, err
	}

	if val.(int64) == -1 {
		return val.(int64), errors.New("This key has no expiration time")
	}

	if val.(int64) == -2 {
		return val.(int64), errors.New("Invalid key")
	}

	return val.(int64), err
}
