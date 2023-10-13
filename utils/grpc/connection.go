package grpc

import (
	"GuTikTok/config"
	"GuTikTok/utils/logging"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"time"
)

func Connect(serviceName string) (conn *grpc.ClientConn) {
	kacp := keepalive.ClientParameters{
		Time:                10 * time.Second, // 如果没有活动，每10秒发送一次ping
		Timeout:             time.Second,      // 等待1秒钟以获取ping的确认回复，否则将认为连接已断开
		PermitWithoutStream: false,            // 即使没有活动的流，也发送ping
	}
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s/%s?wait=15s", config.Conf.Consul.Address, config.Conf.Consul.ConsulAnonymityPrefix+serviceName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithKeepaliveParams(kacp),
	)

	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"service": config.Conf.Consul.ConsulAnonymityPrefix + serviceName,
			"err":     err,
		}).Errorf("Cannot connect to %v service", config.Conf.Consul.ConsulAnonymityPrefix+serviceName)
	}
	return
}
