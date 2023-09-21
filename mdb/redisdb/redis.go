package redisdb

import (
	"context"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"time"
)

// InitRdb 初始化 Redis
func InitRdb(rdb *redis.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("连接redis出错，错误信息：%v", err)
	}
}
