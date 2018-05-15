package session

import (
	"strings"

	"github.com/HiLittleCat/goSeed/config"
	"github.com/HiLittleCat/goSeed/conn"
	redis "gopkg.in/redis.v5"
)

// Expire session expire time
const Expire = 3600

var redisPool = conn.GetRedisPool(config.Default.RedisSession.Name)

// Values session values
type Values map[string]string

// StoreS default session store
var StoreS *Store

// Store redis session store
type Store struct{}

// Create create a session by sid
func (rs *Store) Create(sid string, fields map[string]string) error {
	var err error
	redisPool.Exec(func(c *redis.Client) {
		cmd := c.HMSet(sid, fields)
		err = cmd.Err()
		if err != nil {
			return
		}
		boolCmd := c.Expire(sid, Expire)
		err = boolCmd.Err()
	})
	return err
}

// Get generate a session by sid
func (rs *Store) Get(sid string) (Values, error) {
	var val map[string]string
	var err error
	redisPool.Exec(func(c *redis.Client) {
		cmd := c.HGetAll(sid)
		val, err = cmd.Result()
	})
	return val, err
}

// Delete generate a session by sid
func (rs *Store) Delete(sid string) error {
	var err error
	redisPool.Exec(func(c *redis.Client) {
		cmd := c.Del(sid)
		err = cmd.Err()
	})
	return err
}

// GetFieldValue key value in redis session
func (rs Store) GetFieldValue(sid string, key string) (string, error) {
	var val string
	var err error
	redisPool.Exec(func(c *redis.Client) {
		cmd := c.HGet(sid, key)
		val, err = cmd.Result()
	})
	return val, err
}

// SetFieldValue set key value in redis session
func (rs *Store) SetFieldValue(sid string, key string, value string) error {
	var err error
	redisPool.Exec(func(c *redis.Client) {
		cmd := c.HSet(sid, key, value)
		err = cmd.Err()
		if err != nil {
			return
		}
		boolCmd := c.Expire(sid, Expire)
		err = boolCmd.Err()
	})
	return err
}

// SessionID generate redis session id by uid
func (rs *Store) SessionID(uid string) string {
	sid := uid
	return sid
}

// ToString get session string
func (rs *Store) ToString(values Values) string {
	str := ""
	for k, v := range values {
		str += k + "=" + v + ";"
	}
	return strings.TrimRight(str, ";")
}
