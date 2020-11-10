package my_gocache

import (
	// "log"
	// Import the redigo/redis package.
	"fmt"

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
	} else {
		CSConn = &conn
		return nil
	}
}

// Set method
func (r Redis) Set(key string, value interface{}, expireTime ...int) error {
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
func (r Redis) Delete(key string) error {
	_, err := (*CSConn).Do("DEL", key)

	if err != nil {
		return err
	}

	return nil
}

// Close connection
func (r Redis) Close() {
	(*CSConn).Close()
}

// set expire time on a key
func (r Redis) Expire(key string, expireTime int) (interface{}, error) {
	val, err := (*CSConn).Do("EXPIRE", key, expireTime)
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	} else {
		fmt.Println("val: ", val)
		return val, nil
	}
}
