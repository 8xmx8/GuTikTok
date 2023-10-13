package main

import (
	"GuTikTok/config"
	"GuTikTok/src/extra/profiling"
	"GuTikTok/src/extra/tracing"
	"GuTikTok/src/models"
	"GuTikTok/src/rpc/auth"
	"GuTikTok/src/storage/database"
	"GuTikTok/src/storage/redis"
	"GuTikTok/utils/consul"
	"GuTikTok/utils/logging"
	"GuTikTok/utils/prom"
	"context"
	"github.com/bits-and-blooms/bloom/v3"
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	redis2 "github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"net/http"
	"os"
	"syscall"
)

func main() {
	// 设置分布式追踪提供者
	tp, err := tracing.SetTraceProvider(config.AuthRpcServerName)

	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Panicf("Error to set the trace")
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			logging.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorf("Error to set the trace")
		}
	}()

	// 配置 Pyroscope
	profiling.InitPyroscope("GuGoTik.AuthService")

	// 初始化性能分析工具
	log := logging.LogService(config.AuthRpcServerName)
	// 创建 TCP 监听器
	lis, err := net.Listen("tcp", config.Conf.Server.Address+config.AuthRpcServerPort)

	if err != nil {
		log.Panicf("Rpc %s listen happens error: %v", config.AuthRpcServerName, err)
	}
	// gRPC 服务器中收集和暴露指标
	srvMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	// 注册到 Prometheus 客户端注册器
	reg := prom.Client
	reg.MustRegister(srvMetrics)
	// Create a new Bloom filter with a target false positive rate of 0.1%
	BloomFilter = bloom.NewWithEstimates(10000000, 0.001) // assuming we have 1 million users

	// Initialize BloomFilter from database
	var users []models.User
	userNamesResult := database.Client.WithContext(context.Background()).Select("name").Find(&users)

	if userNamesResult.Error != nil {
		log.Panicf("Getting user names from databse happens error: %s", userNamesResult.Error)
		panic(userNamesResult.Error)
	}
	for _, u := range users {
		BloomFilter.AddString(u.Name)
	}

	// Create a go routine to receive redis message and add it to BloomFilter
	go func() {
		pubSub := redis.Client.Subscribe(context.Background(), config.BloomRedisChannel)
		defer func(pubSub *redis2.PubSub) {
			err := pubSub.Close()
			if err != nil {
				log.Panicf("Closing redis pubsub happend error: %s", err)
			}
		}(pubSub)

		_, err := pubSub.ReceiveMessage(context.Background())
		if err != nil {
			log.Panicf("Reveiving message from redis happens error: %s", err)
			panic(err)
		}

		ch := pubSub.Channel()
		for msg := range ch {
			log.Infof("Add user name to BloomFilter: %s", msg.Payload)
			BloomFilter.AddString(msg.Payload)
		}
	}()

	s := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.ChainUnaryInterceptor(srvMetrics.UnaryServerInterceptor(grpcprom.WithExemplarFromContext(prom.ExtractContext))),
		grpc.ChainStreamInterceptor(srvMetrics.StreamServerInterceptor(grpcprom.WithExemplarFromContext(prom.ExtractContext))),
	)
	if err := consul.RegisterConsul(config.AuthRpcServerName, config.AuthRpcServerPort); err != nil {
		log.Panicf("Rpc %s register consul happens error for: %v", config.AuthRpcServerName, err)
	}
	log.Infof("Rpc %s is running at %s now", config.AuthRpcServerName, config.AuthRpcServerPort)

	var srv AuthServiceImpl
	auth.RegisterAuthServiceServer(s, srv)
	// 将健康服务器注册到 gRPC 服务器
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	srv.New()
	srvMetrics.InitializeMetrics(s)

	g := &run.Group{}
	g.Add(func() error {
		return s.Serve(lis)
	}, func(err error) {
		s.GracefulStop()
		s.Stop()
		log.Errorf("Rpc %s listen happens error for: %v", config.AuthRpcServerName, err)
	})

	httpSrv := &http.Server{Addr: config.Conf.Server.Address + config.Metrics}
	g.Add(func() error {
		m := http.NewServeMux()
		m.Handle("/metrics", promhttp.HandlerFor(
			reg,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			},
		))
		httpSrv.Handler = m
		log.Infof("Promethus now running")
		return httpSrv.ListenAndServe()
	}, func(error) {
		if err := httpSrv.Close(); err != nil {
			log.Errorf("Prometheus %s listen happens error for: %v", config.AuthRpcServerName, err)
		}
	})

	g.Add(run.SignalHandler(context.Background(), syscall.SIGINT, syscall.SIGTERM))

	if err := g.Run(); err != nil {
		log.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("Error when runing http server")
		os.Exit(1)
	}
}
