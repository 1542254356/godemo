package command

import (
	"testing"

	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	conn, err := GetConn()
	require.NoError(t, err)
	defer conn.Close()

	k := "h"
	reply, err := conn.Do("DEL", k)
	require.NoError(t, err)

	reply, err = conn.Do("HSET", k, "blog", "hjwblog.com")
	require.NoError(t, err)
	assert.Equal(t, int64(1), reply)

	blog, err := redis.String(conn.Do("HGET", k, "blog"))
	require.NoError(t, err)
	assert.Equal(t, "hjwblog.com", blog)

	m, err := redis.StringMap(conn.Do("HGETALL", k))
	require.NoError(t, err)
	assert.Equal(t, map[string]string{
		"blog": "hjwblog.com",
	}, m)

	reply, err = conn.Do("HDEL", k, "blog")
	require.NoError(t, err)
	assert.Equal(t, int64(1), reply)
}
