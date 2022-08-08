package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisTest struct {
	conn       *redis.Client
	subscriber *redis.PubSub
}

func (r *redisTest) receivedPoint() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if _, err := r.subscriber.ReceiveMessage(ctx); err != nil {
		return err
	}

	return nil
}

func initTestRedis(conf map[string]string) redisTest {
	result := redisTest{}

	result.conn = redis.NewClient(&redis.Options{Addr: conf["server"]})
	result.subscriber = result.conn.Subscribe(context.Background(), conf["queue"])

	return result
}
