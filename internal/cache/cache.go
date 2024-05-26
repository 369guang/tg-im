package cache

var RDS *RedisClient

func InitCache(host, password string, port, db int) {
	RDS = NewRedisClient(host, password, port, db)
}
