package command

import (
	"os"

	"github.com/garyburd/redigo/redis"
)

// GetConn Get redis Conn
func GetConn() (redis.Conn, error) {
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWD")
	if redisAddr == "" {
		redisAddr = "127.0.0.1:6379"
	}
	return redis.Dial("tcp", redisAddr, redis.DialPassword(redisPassword))
}
