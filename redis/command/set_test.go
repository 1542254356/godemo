package command

import (
	"github.com/garyburd/redigo/redis"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	conn, err := GetConn()
	require.NoError(t, err)
	defer conn.Close()

	k := "s"
	reply, err := conn.Do("DEL", k)
	require.NoError(t, err)

	reply, err = conn.Do("SADD", k, 1, 2, 3)
	require.NoError(t, err)
	assert.Equal(t, int64(3), reply)

	reply, err = conn.Do("SADD", k, 1, 2, 3, 4)
	require.NoError(t, err)
	assert.Equal(t, int64(1), reply)
	
	isMember, err := redis.Bool(conn.Do("SISMEMBER", k, 4))
	require.NoError(t, err)
	assert.Equal(t, true,isMember)
	
	reply, err = conn.Do("SREM", k, 4)
	require.NoError(t, err)
	assert.Equal(t, int64(1),reply)
	
	isMember, err = redis.Bool(conn.Do("SISMEMBER", k, 4))
	require.NoError(t, err)
	assert.Equal(t, false,isMember)
	
	setMembers, err := redis.Ints(conn.Do("SMEMBERS", k))
	require.NoError(t, err)
	assert.Equal(t,[]int{1,2,3},setMembers)
}
