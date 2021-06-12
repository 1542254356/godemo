package command

import (
	"github.com/garyburd/redigo/redis"
	"testing"
	"time"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMulti(t *testing.T) {
	k1, k2 := "k1", "k2"

	conn, err := GetConn()
	require.NoError(t, err)
	defer conn.Close()

	reply, err := conn.Do("DEL", k1, k2)
	require.NoError(t, err)

	reply, err = conn.Do("SET", k1, 0)
	require.NoError(t, err)
	assert.Equal(t, "OK", reply)

	reply, err = conn.Do("SET", k2, 0)
	require.NoError(t, err)
	assert.Equal(t, "OK", reply)

	const N = 100
	for i := 0; i < N; i++ {
		go Inc(t, k1, k2)
	}
	
	time.Sleep(time.Millisecond *500)
	
	for i := 0; i < 100; i++ {
		list, err := redis.Ints(conn.Do("MGET", k1,k2))
		require.NoError(t, err)
		t.Log(list[0],list[1])
		
		assert.Equal(t, list[0], list[1])
	}
	
	time.Sleep(time.Second * 5)
}

func Inc(t *testing.T, k1, k2 string) {
	conn, err := GetConn()
	require.NoError(t, err)
	defer conn.Close()

	reply, err := conn.Do("MULTI")
	require.NoError(t, err)
	assert.Equal(t, "OK", reply)

	reply, err = conn.Do("Incr", k1)
	require.NoError(t, err)
	assert.Equal(t, "QUEUED", reply)

	time.Sleep(time.Millisecond * 500)

	reply, err = conn.Do("Incr", k2)
	require.NoError(t, err)
	assert.Equal(t, "QUEUED", reply)

	reply, err = conn.Do("EXEC")
	require.NoError(t, err)
	t.Log(reply)
}
