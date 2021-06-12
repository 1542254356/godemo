package command

import (
	"testing"

	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPipeline(t *testing.T) {
	conn, err := GetConn()
	require.NoError(t, err)
	defer conn.Close()
	k := "p"

	go func() {
		err2 := conn.Send("SET", k, "1")
		require.NoError(t, err2)

		err2 = conn.Send("GET", k)
		require.NoError(t, err2)

		err2 = conn.Send("GET", k)
		require.NoError(t, err2)

		err2 = conn.Flush()
		require.NoError(t, err2)
	}()

	reply, err := conn.Receive()
	require.NoError(t, err)
	assert.Equal(t, "OK", reply)

	v, err := redis.Int(conn.Receive())
	require.NoError(t, err)
	assert.Equal(t, 1, v)

	v, err = redis.Int(conn.Receive())
	require.NoError(t, err)
	assert.Equal(t, 1, v)
}
