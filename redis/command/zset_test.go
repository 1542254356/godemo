package command

import (
	"testing"

	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestZset(t *testing.T) {
	conn, err := GetConn()
	require.NoError(t, err)
	defer conn.Close()

	k := "z"
	reply, err := conn.Do("DEL", k)
	require.NoError(t, err)

	reply, err = conn.Do("ZADD", k, 8, "8", 7, "7")
	require.NoError(t, err)
	assert.Equal(t, int64(2), reply)

	m, err := redis.IntMap(conn.Do("ZRANGEBYSCORE", k, 0, 100, "withScores"))
	require.NoError(t, err)
	assert.Equal(t, map[string]int{
		"8": 8,
		"7": 7,
	}, m)

	reply, err = conn.Do("ZREM", k, "8")
	require.NoError(t, err)
	assert.Equal(t, int64(1), reply)
	
	m, err = redis.IntMap(conn.Do("ZRANGEBYSCORE", k, 0, 100, "withScores"))
	require.NoError(t, err)
	assert.Equal(t, map[string]int{
		"7": 7,
	}, m)

}
