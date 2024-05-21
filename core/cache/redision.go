package cache

import (
	"context"
	"fmt"
	"github.com/369guang/tg-im/core/logs"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

var ctx = context.Background()

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(host, password string, port, db int) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       db,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		logs.Logger.Error("redis ping error", zap.Error(err))
		panic(err)
	}

	return &RedisClient{Client: client}
}

// Set 设置一个键值对到redis
func (r *RedisClient) Set(key string, value interface{}) error {
	return r.Client.Set(ctx, key, value, 0).Err()
}

// Get 从redis获取一个键值对
func (r *RedisClient) Get(key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

// Del 从redis删除一个键值对
func (r *RedisClient) Del(key string) error {
	return r.Client.Del(ctx, key).Err()
}

// Exists 判断一个键是否存在
func (r *RedisClient) Exists(key string) (int64, error) {
	return r.Client.Exists(ctx, key).Result()
}

// Expire 设置一个键的过期时间
func (r *RedisClient) Expire(key string, expiration time.Duration) error {
	return r.Client.Expire(ctx, key, expiration).Err()
}

// SetHash 设置一个hash键值对到redis
func (r *RedisClient) SetHash(key string, field string, value interface{}) error {
	return r.Client.HSet(ctx, key, field, value).Err()
}

// GetHash 从redis获取一个hash键值对
func (r *RedisClient) GetHash(key string, field string) (string, error) {
	return r.Client.HGet(ctx, key, field).Result()
}

// DelHash 从redis删除一个hash键值对
func (r *RedisClient) DelHash(key string, field string) error {
	return r.Client.HDel(ctx, key, field).Err()
}

// ExistsHash 判断一个hash键是否存在
func (r *RedisClient) ExistsHash(key string, field string) (bool, error) {
	return r.Client.HExists(ctx, key, field).Result()
}

// Publish 发布消息
func (r *RedisClient) Publish(channel string, message interface{}) error {
	return r.Client.Publish(ctx, channel, message).Err()
}
