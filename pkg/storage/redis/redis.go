package redis

import (
	"context"
	"fmt"

	c "github.com/mrceylan/go-url-shortener/pkg/configuration"

	"github.com/go-redis/redis/v8"
)

type Connection struct {
	Client *redis.Client
}

var (
	ctx = context.Background()
)

func RedisInit() *Connection {
	defer fmt.Println("Redis client connected")
	return &Connection{redis.NewClient(&redis.Options{
		Addr:     c.Config.Redis.Connection,
		Password: c.Config.Redis.Password,
		DB:       c.Config.Redis.Db,
	})}

}
