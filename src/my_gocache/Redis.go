package my_gocache

import (
	// "log"
	// Import the redigo/redis package.
	"github.com/gomodule/redigo/redis"
)

var (
	// CSConn: Connector
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
func (r Redis) Set(key string, value interface{}) error {
	_, err := (*CSConn).Do("SET", key, value)

	if err != nil {
		return err
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
