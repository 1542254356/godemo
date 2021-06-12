package command

import (
	"testing"

	"github.com/garyburd/redigo/redis"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMgetMset(t *testing.T) {
	conn, err := GetConn()
	require.NoError(t, err)
	defer conn.Close()

	reply, err := conn.Do("MSET", "name", "hjw", "blog", "hjwblog.com")
	assert.NoError(t, err)
	assert.Equal(t, "OK", reply)

	values, err := redis.Strings(conn.Do("MGET", "name", "blog"))
	require.NoError(t, err)
	assert.Equal(t, []string{"hjw", "hjwblog.com"}, values)
}
