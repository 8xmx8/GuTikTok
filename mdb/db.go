package mdb

import (
	"GuTikTok/config"
	"GuTikTok/mdb/model"
	"fmt"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	stdlog "log"
	"os"
	"time"
)

var DB *gorm.DB //操作数据库入口

func InitLog() {

	logConf := config.Conf.Log
	if logConf.Enable {

		log.SetLevel(log.InfoLevel)
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

func InitDb() {
	var dialector gorm.Dialector
	var err error

	database := config.Conf.MySQL
	dialector = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		database.UserName, database.Password, database.Host, database.Port, database.Database, database.Charset))

	logLevel := logger.Info
	DB, err = gorm.Open(dialector, &gorm.Config{
		//日志配置后期要更改
		Logger: logger.New(
			stdlog.New(log.StandardLogger().Out, "\r\n", stdlog.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second, // 设置慢查询阈值
				LogLevel:                  logLevel,    // 设置日志级别
				IgnoreRecordNotFoundError: true,        // 忽略记录未找到的错误
				Colorful:                  true,        // 启用彩色日志输出
			},
		),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //表名以单数形式命名
		},
		TranslateError: true, // 启用错误翻译功能
	})
	if err != nil {
		log.Fatalf("无法连接到数据库:%s", err.Error())
	}
	model.Init(DB)
	log.Info("初始化 Database 成功!")
}
