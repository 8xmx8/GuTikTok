package mdb

import (
	"GuTikTok/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"time"
)

var rdb *redis.Client //操作redis入口

// InitRdb 初始化 Redis
func InitRdb() {
	log.Info("开始初始化 Redis 服务!")
	rconf := config.Conf.Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", rconf.Host, rconf.Port),
		Password: rconf.Password,
		DB:       rconf.Db,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("连接redis出错，错误信息：%v", err)
	}
	log.Info("初始化 Redis 成功!")
}