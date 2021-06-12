package command

import (
	"testing"

	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSet(t *testing.T) {
	conn, err := GetConn()
	require.NoError(t, err)
	defer conn.Close()

	v := "hjwblog.com"
	k := "name"
	reply, err := conn.Do("SET", k, v)
	require.NoError(t, err)
	assert.Equal(t, "OK", reply)
	t.Log(reply)

	name, err := redis.String(conn.Do("GET", k))
	require.NoError(t, err)
	t.Log(name)
	assert.Equal(t, v, name)

	reply, err = conn.Do("DEL", k)
	require.NoError(t, err)
	t.Log(reply)
	assert.Equal(t, int64(1), reply)

	reply, err = conn.Do("GET", k)
	require.NoError(t, err)
	assert.Equal(t, nil, reply)
	t.Log(reply)
}
