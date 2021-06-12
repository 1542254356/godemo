package command

import "github.com/garyburd/redigo/redis"

// GetPool Get redis connection Pool
func GetPool() redis.Pool {
	return redis.Pool{
		Dial: GetConn,
		MaxIdle:         10,
		MaxActive:       100,
		IdleTimeout:     600,
	}
}
