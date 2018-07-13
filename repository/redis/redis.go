package redis

import (
	"os"

	"github.com/gomodule/redigo/redis"
)

// Redis struct which is responsible for maintaining redis functionality
type Redis struct {
	conn redis.Conn
}

// New returns pointer to brand new Redis struct
func New() *Redis {
	return &Redis{}
}

// Connect to db
func (r *Redis) Connect() (err error) {
	r.conn, err = redis.DialURL(os.Getenv("REDIS_URL"))
	return
}

// Close function closes redis
func (r Redis) Close() {
	r.conn.Close()
}

// Fetch function fetches data from db
func (r Redis) Fetch(uuid string) ([]byte, error) {
	n, err := r.conn.Do("JSON.GET", uuid)
	if n == nil {
		n = []byte{}
	}
	return n.([]byte), err
}

// Put function persists data into redis
func (r Redis) Put(key string, data []byte) error {
	_, err := r.conn.Do("JSON.SET", key, ".", data)
	return err
}
