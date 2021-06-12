package command

import (
	"testing"
	"time"
	
	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/require"
)

func TestPubSub(t *testing.T) {
	ticker := time.NewTicker(time.Millisecond * 500)

	go Sub(t)
	go Publish(t, ticker)
	
	time.Sleep(time.Second * 5)
	ticker.Stop()
}

func Publish(t *testing.T,ticker *time.Ticker) {
	conn, err := GetConn()
	require.NoError(t, err)
	defer conn.Close()
	
	for range ticker.C {
		reply, err := conn.Do("PUBLISH", "c", time.Now())
		require.NoError(t, err)
		t.Logf("sub count %d",reply)
	}
}

func Sub(t *testing.T) {
	conn, err := GetConn()
	require.NoError(t, err)
	defer conn.Close()

	subConn := redis.PubSubConn{Conn: conn}
	err = subConn.Subscribe("c")
	require.NoError(t, err)

	for {
		switch v := subConn.Receive().(type) {
		case redis.Message:
			t.Logf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			t.Logf("%s: %s %d", v.Channel, v.Kind, v.Count)
		case error:
			require.NoError(t, err)
		}
	}
}
