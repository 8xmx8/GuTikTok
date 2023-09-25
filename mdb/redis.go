package mdb

import (
	"GuTikTok/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/extra/redisotel/v9"
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
	/*
		redisotel.InstrumentTracing 用于启用 Redis 客户端的追踪功能，
		它会自动将 Redis 操作记录到追踪系统中，以便进行性能分析和故障排查。
		这可以帮助你了解每个 Redis 操作的执行时间、调用关系和可能的问题点。

		redisotel.InstrumentMetrics 用于启用 Redis 客户端的指标监控功能，
		它会自动收集有关 Redis 操作的各种指标信息，如请求计数、错误计数、响应时间等。
		这可以帮助你实时监控 Redis 的性能和健康状况，并进行适当的调整和优化。
	*/
	if err := redisotel.InstrumentTracing(rdb); err != nil {
		panic(err)
	}

	if err := redisotel.InstrumentMetrics(rdb); err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("连接redis出错，错误信息：%v", err)
	}
	log.Info("初始化 Redis 成功!")
}
