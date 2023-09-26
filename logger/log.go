package logger

import (
	"GuTikTok/config"
	"fmt"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func init() {

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

	// 配置日志格式
	formatter := log.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           "2006-01-02 15:04:05",
		FullTimestamp:             true,
		DisableQuote:              true,
	}
	log.SetFormatter(&formatter)

	log.Info("初始化 logrus 成功!")
}
