package command

import (
	"testing"

	"github.com/garyburd/redigo/redis"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	conn, err := GetConn()
	require.NoError(t, err)
	defer conn.Close()

	k := "l"
	reply, err := conn.Do("DEL", k)
	require.NoError(t, err)

	reply, err = conn.Do("LPUSH", k, 2, 1, 0)
	require.NoError(t, err)
	assert.Equal(t, int64(3), reply)

	list, err := redis.Ints(conn.Do("LRANGE", k, 0, -1))
	require.NoError(t, err)
	assert.Equal(t, []int{0, 1, 2}, list)

	n, err := redis.Int(conn.Do("LPOP", k))
	require.NoError(t, err)
	assert.Equal(t, 0, n)

	list, err = redis.Ints(conn.Do("LRANGE", k, 0, -1))
	require.NoError(t, err)
	assert.Equal(t, []int{1, 2}, list)

	reply, err = conn.Do("RPUSH", k, 3, 4, 5)
	require.NoError(t, err)
	assert.Equal(t, int64(5), reply)

	n, err = redis.Int(conn.Do("RPOP", k))
	require.NoError(t, err)
	assert.Equal(t, 5, n)

	list, err = redis.Ints(conn.Do("LRANGE", k, 0, -1))
	require.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3, 4}, list)

	n, err = redis.Int(conn.Do("LINDEX", k, 0))
	require.NoError(t, err)
	assert.Equal(t, 1, n)
}
