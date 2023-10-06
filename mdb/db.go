package mdb

import (
	"GuTikTok/config"
	"GuTikTok/mdb/model"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"gorm.io/plugin/opentelemetry/tracing"
	stdlog "log"
	"strings"
	"time"
)

var DB *gorm.DB //操作数据库入口

func init() {
	log.Info("开始初始化 Database !")
	var dialector gorm.Dialector
	var logLevel logger.LogLevel
	var err error

	database := config.Conf.MySQL
	dialector = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		database.UserName, database.Password, database.Host, database.Port, database.Database, database.Charset))

	switch strings.ToLower(database.LogLevel) {
	case "silent":
		logLevel = logger.Silent
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	case "info":
		logLevel = logger.Info
	default:
		log.Fatalf("数据库日志级别不正确，可用: [silent,error,warn,info]")
	}

	DB, err = gorm.Open(dialector, &gorm.Config{
		PrepareStmt: true, //启用预编译sql
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
	} else {
		model.Init(DB)
	}

	if config.Conf.MysqlReplica.MySQLReplicaState == "enable" {
		var replicas []gorm.Dialector

		for _, addr := range strings.Split(config.Conf.MysqlReplica.MySQLReplicaAddress, ",") {
			// 这里可能需要根据 MySQL 连接字符串的格式进行修改，以适应您的实际情况
			replicaDSN := fmt.Sprintf("%s:%s@tcp(%s)/%s",
				config.Conf.MysqlReplica.MySQLReplicaUsername,
				config.Conf.MysqlReplica.MySQLReplicaPassword,
				addr,
				config.Conf.MySQL.Database)

			// 使用 MySQL 数据源名称连接到副本数据库
			replicas = append(replicas, mysql.Open(replicaDSN))
		}

		// 注册数据库副本和负载均衡策略
		err := DB.Use(dbresolver.Register(dbresolver.Config{
			Replicas: replicas,
			Policy:   dbresolver.RandomPolicy{},
		}))
		if err != nil {
			panic(err)
		}
	}

	// 获取数据库连接对象
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}

	// 配置数据库连接池的参数
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(24 * time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Hour)

	// 添加追踪插件
	if err := DB.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}
	log.Info("初始化 Database 成功!")
}
