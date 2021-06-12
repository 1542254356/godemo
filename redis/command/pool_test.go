package command

import (
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPool(t *testing.T) {
	pool := GetPool()

	reply, err := pool.Get().Do("Set", "k", "v")
	require.NoError(t, err)
	assert.Equal(t, "OK", reply)

	const N = 200
	for i := 0; i < N; i++ {
		go SetGet(t, pool.Get())
	}

	time.Sleep(time.Second * 5)
}

func SetGet(t *testing.T, conn redis.Conn) {
	defer conn.Close()
	n, err := redis.String(conn.Do("Get", "k"))
	require.NoError(t, err)
	t.Logf("conn:%p,%s",conn,n)
}
