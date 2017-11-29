package conn

const (
	RedisBosh = "bosh"
	SessionDB = 0
)

func init() {
	redisPools = make(map[string]*RedisPool)
}

var redisPools map[string]*RedisPool

func RedisSet(key string, p *RedisPool) {
	redisPools[key] = p
}

func GetRedisPool(key string) *RedisPool {
	return redisPools[key]
}
