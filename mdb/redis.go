package mdb

import (
	"GuTikTok/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"time"
)

var rdb redis.UniversalClient //操作redis入口
const redisName = ""

// InitRdb 初始化 Redis
func InitRdb() {
	log.Info("开始初始化 Redis 服务!")
	rconf := config.Conf.Redis
	redis_addr := fmt.Sprintf("%s:%d", rconf.Host, rconf.Port)
	rdb = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{
			redis_addr,
		},
		Password:   rconf.Password,
		DB:         rconf.Db,
		MasterName: redisName,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("连接redis出错，错误信息：%v", err)
	}
	log.Info("初始化 Redis 成功!")
}
