package command

import (
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExpire(t *testing.T) {
	conn, err := GetConn()
	require.NoError(t, err)
	defer conn.Close()

	k := "expire_k"
	reply, err := conn.Do("SET", k, "v")
	require.NoError(t, err)
	assert.Equal(t, "OK", reply)

	reply, err = conn.Do("EXPIRE", k, 1)
	require.NoError(t, err)
	assert.Equal(t, int64(1), reply)

	v, err := redis.String(conn.Do("GET", k))
	require.NoError(t, err)
	t.Log(v)

	time.Sleep(time.Second)
	v, err = redis.String(conn.Do("GET", k))
	assert.Error(t, err)
	t.Log(err)
}
