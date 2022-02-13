package redis

import (
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
	addr   string
}

var redisInstance *Redis

func StaticInstance(addr string) *Redis {

	if redisInstance == nil {
		redisInstance = &Redis{
			client: redis.NewClient(&redis.Options{
				Addr: addr,
			}),
			addr: addr,
		}
	}

	return redisInstance
}

func (r *Redis) Client() *redis.Client {
	return r.client
}
