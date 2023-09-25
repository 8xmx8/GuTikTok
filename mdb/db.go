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
	stdlog "log"
	"time"
)

var DB *gorm.DB //操作数据库入口

func init() {
	var dialector gorm.Dialector
	var err error

	database := config.Conf.MySQL
	dialector = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		database.UserName, database.Password, database.Host, database.Port, database.Database, database.Charset))

	logLevel := logger.Info
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
	}
	model.Init(DB)
	log.Info("初始化 Database 成功!")
}
