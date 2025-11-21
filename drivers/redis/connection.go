package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
	"dm_loanservice/drivers/goconf"
	"time"
)

var (
	connection *redis.Client
	mutex      sync.Mutex
)

func GetConnection(ctx context.Context) *redis.Client {
	if connection == nil {
		mutex.Lock()
		defer mutex.Unlock()
		connection = newConnection(ctx)
	}

	return connection
}

func newConnection(ctx context.Context) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         goconf.Config().GetString("redis.address"),
		Password:     goconf.Config().GetString("redis.password"),
		PoolTimeout:  2 * time.Second,
		IdleTimeout:  2 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Panicf("Got an error while connecting redis server, error: %s", err)
	}

	return rdb
}
