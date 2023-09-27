package logging

import (
	"GuTikTok/config"
	"fmt"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"io"
	"os"
)

var hostname string

// 初始化函数，在包被导入时执行
func init() {

	// 获取主机名
	hostname, _ = os.Hostname() //acer_yjy

	// 配置日志格式
	formatter := log.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           "2006-01-02 15:04:05",
		FullTimestamp:             true,
		DisableQuote:              true,
	}

	log.SetFormatter(&formatter)
	log.AddHook(logTraceHook{})

	logConf := config.Conf.Log
	if logConf.Enable {
		level, err := log.ParseLevel(logConf.Level)
		if err != nil {
			panic(fmt.Sprintf("日志级别不正确，可用: [panic,fatal,error,warn,info,debug,trace],%v", err))
		}
		log.SetLevel(level)
		log.SetReportCaller(true)
		var w io.Writer = &lumberjack.Logger{
			Filename:   logConf.Name,
			MaxSize:    logConf.MaxSize,
			MaxBackups: logConf.MaxBackups,
			MaxAge:     logConf.MaxAge,
			Compress:   logConf.Compress,
		}
		w = io.MultiWriter(os.Stdout, w)

		log.SetOutput(w)
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

	// 创建Logger实例，并设置预定义的字段值
	Logger = log.WithFields(log.Fields{
		"Tied":     "NONE",
		"Hostname": hostname,
		"PodIP":    config.Conf.Server.Address,
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

	if config.Conf.Log.Enable {
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
	span.SetAttributes(attribute.String("podIP", config.Conf.Server.Address))
}
