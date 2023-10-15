package tracing

import (
	"GuTikTok/config"
	"GuTikTok/utils/logging"
	"context"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	trace2 "go.opentelemetry.io/otel/trace"
)

var Tracer trace2.Tracer // 定义一个名为 Tracer 的 trace2.Tracer 类型的变量

func SetTraceProvider(name string) (*trace.TracerProvider, error) {
	// 创建 OpenTelemetry 的 HTTP 客户端
	client := otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint(config.Conf.Tracers.Addr), // 将跟踪数据发送到指定的 IP 地址上的服务
		otlptracehttp.WithInsecure(),                         // 允许不安全的连接
	)

	// 创建 OpenTelemetry 的导出器（exporter）
	exporter, err := otlptrace.New(context.Background(), client)
	if err != nil {
		// 如果创建导出器时发生错误，则记录日志并返回错误
		logging.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("Can not init otel!")
		return nil, err
	}

	var sampler trace.Sampler
	// 根据配置文件中的 Tracers.OtelState 参数设置采样器
	if config.Conf.Tracers.OtelState == "disable" {
		// 如果采样状态为 "disable"，则使用 trace.NeverSample() 禁用采样
		sampler = trace.NeverSample()
	} else {
		// 如果采样状态不为 "disable"，则根据配置文件中的 Tracers.OtelSampler 参数使用 trace.TraceIDRatioBased() 创建基于跟踪 ID 比率的采样器
		sampler = trace.TraceIDRatioBased(config.Conf.Tracers.OtelSampler)
	}

	// 创建 OpenTelemetry 的跟踪提供者（tracer provider），并配置导出器、资源和采样器参数
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter), // 使用导出器作为批处理器
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(name),
			),
		),
		trace.WithSampler(sampler), // 设置采样器
	)

	// 设置全局的 OpenTelemetry 跟踪提供者
	otel.SetTracerProvider(tp)

	// 设置 OpenTelemetry 的文本传播器，用于在跨系统和服务之间传递跟踪上下文信息
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// 为指定名称创建 OpenTelemetry 的跟踪器（tracer）
	Tracer = otel.Tracer(name)

	return tp, nil // 返回跟踪器提供者和 nil 作为错误
}
