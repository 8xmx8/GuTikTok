package logging

import (
	"GuTikTok/config"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"io"
	"os"
	"path"
	"strings"
)

var hostname string

// 初始化函数，在包被导入时执行
func init() {
	InitDBLog()
	// 获取主机名
	hostname, _ = os.Hostname() //acer_yjy

	// 根据配置的日志级别设置日志记录器的级别
	switch strings.ToUpper(config.Conf.Log.Level) {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN", "WARNING":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	}

	// 设置日志文件的路径和创建文件夹
	filePath := path.Join("data", "log", "gutiktok.log")
	dir := path.Dir(filePath)
	if err := os.MkdirAll(dir, os.FileMode(0755)); err != nil {
		panic(err)
	}

	// 打开日志文件，并设置日志记录器的格式和输出位置
	/*
		O_RDONLY 打开只读文件
		O_WRONLY 打开只写文件
		O_RDWR 打开既可以读取又可以写入文件
		O_APPEND 写入文件时将数据追加到文件尾部
		O_CREATE 如果文件不存在，则创建一个新的文件
		O_TRUNC 表示如果文件存在，则截断文件到零长度
		0o666：表示文件权限的八进制数。0o666 表示文件所有者、所属组和其他用户都具有读写权限。
	*/
	/*
		在八进制表示法中，0o 前缀表示八进制数。数字 766 对应了文件权限 rw-rw-rw-。
		7 表示所有者（owner）具有读取、写入和执行权限。
		6 表示所属组（group）具有读取和写入权限。
		6 表示其他用户（others）具有读取和写入权限。
	*/
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	log.SetFormatter(&log.JSONFormatter{})
	log.AddHook(logTraceHook{})
	log.SetOutput(io.MultiWriter(f, os.Stdout))

	// 创建Logger实例，并设置预定义的字段值
	Logger = log.WithFields(log.Fields{
		"Tied":     "",
		"Hostname": hostname,
		"PodIP":    "",
	})
}

// 自定义的logrus钩子实现，用于向日志条目添加OpenTelemetry跟踪信息
type logTraceHook struct{}

// Levels方法指定该钩子应应用于所有日志级别
func (t logTraceHook) Levels() []log.Level { return log.AllLevels }

// Fire方法在进行日志记录时调用，将OpenTelemetry跟踪信息添加到日志条目的数据中
func (t logTraceHook) Fire(entry *log.Entry) error {
	ctx := entry.Context
	if ctx == nil {
		return nil
	}

	span := trace.SpanFromContext(ctx)
	//if !span.IsRecording() {
	//	return nil
	//}

	sCtx := span.SpanContext()
	if sCtx.HasTraceID() {
		entry.Data["trace_id"] = sCtx.TraceID().String()
	}
	if sCtx.HasSpanID() {
		entry.Data["span_id"] = sCtx.SpanID().String()
	}

	if config.Conf.LoggerWithTraceState == "enable" {
		attrs := make([]attribute.KeyValue, 0)
		logSeverityKey := attribute.Key("log.severity")
		logMessageKey := attribute.Key("log.message")
		attrs = append(attrs, logSeverityKey.String(entry.Level.String()))
		attrs = append(attrs, logMessageKey.String(entry.Message))
		for key, value := range entry.Data {
			fields := attribute.Key(fmt.Sprintf("log.fields.%s", key))
			attrs = append(attrs, fields.String(fmt.Sprintf("%v", value)))
		}
		span.AddEvent("log", trace.WithAttributes(attrs...))
		if entry.Level <= log.ErrorLevel {
			span.SetStatus(codes.Error, entry.Message)
		}
	}
	return nil
}

var Logger *log.Entry

// LogService函数返回一个具有指定"Service"字段的新日志条目
func LogService(name string) *log.Entry {
	return Logger.WithFields(log.Fields{
		"Service": name,
	})
}

// 设置OpenTelemetry跟踪的Span为错误状态，并记录错误
func SetSpanError(span trace.Span, err error) {
	span.RecordError(err)
	span.SetStatus(codes.Error, "Internal Error")
}

// 设置OpenTelemetry跟踪的Span为错误状态，并记录错误并提供描述信息
func SetSpanErrorWithDesc(span trace.Span, err error, desc string) {
	span.RecordError(err)
	span.SetStatus(codes.Error, desc)
}
func SetSpanWithHostname(span trace.Span) {
	span.SetAttributes(attribute.String("hostname", hostname))
	span.SetAttributes(attribute.String("podIP", "localhost"))
}
