package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func TestRedisCluster(t *testing.T) {
	// rdb := redis.NewClusterClient(&redis.ClusterOptions{
	// 	Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},

	// 	// To route commands by latency or randomly, enable one of the following.
	// 	// RouteByLatency: true,
	// 	// RouteRandomly: true,
	// })

	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}