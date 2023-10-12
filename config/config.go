package config

import (
	"GuTikTok/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var Conf *Config

type Config struct {
	JwtSecret    string        `yaml:"JwtSecret"`
	Server       *Server       `yaml:"server"`
	Consul       *Consul       `yaml:"consul"`
	Tracers      *Tracers      `yaml:"tracers"`
	MySQL        *Mysql        `yaml:"mysql"`
	MysqlReplica *MysqlReplica `yaml:"MysqlReplica"`
	Log          *Log          `yaml:"log"`
	Redis        *Redis        `yaml:"redis"`
	Pyroscope    *Pyroscope    `yaml:"pyroscope"`
}
type MysqlReplica struct {
	MySQLReplicaState    string `yaml:"MySQLReplicaState"`
	MySQLReplicaAddress  string `yaml:"MySQLReplicaAddress"`
	MySQLReplicaUsername string `yaml:"MySQLReplicaUsername"`
	MySQLReplicaPassword string `yaml:"MySQLReplicaPassword"`
}
type Server struct {
	Address  string `yaml:"address"`
	Https    bool   `yaml:"https"`    //是否启用https
	CertFile string `yaml:"certFile"` // 证书路径
	KeyFile  string `yaml:"keyFile"`  // 证书路径
}
type Consul struct {
	Address               string `yaml:"address"`
	ConsulAnonymityPrefix string `yaml:"consulAnonymityPrefix"`
}
type Tracers struct {
	OtelState   string  `yaml:"OtelState"`   //diable->禁用OpenTelemetry采样功能
	OtelSampler float64 `yaml:"OtelSampler"` //采样比例->0.01
}
type Pyroscope struct { //分析器
	PyroscopeState string `yaml:"PyroscopeState"`
	PyroscopeAddr  string `yaml:"PyroscopeAddr"`
}
type Mysql struct {
	LogLevel string `yaml:"logLevel"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}
type Log struct {
	Enable               bool   `yaml:"enable"` // 是否启用日志
	LoggerWithTraceState string `yaml:"loggerWithTraceState"`
	Level                string `yaml:"info"`       // 日志等级，可用 panic,fatal,error,warn,info,debug,trace
	Name                 string `yaml:"name"`       // 日志文件名
	MaxSize              int    `yaml:"MaxSize"`    // 日志最大大小
	MaxBackups           int    `yaml:"MaxBackups"` // 日志最大备份数
	MaxAge               int    `yaml:"MaxAge"`     // 日志最长时间
	Compress             bool   `yaml:"compress"`   // 日志是否压缩
}
type Redis struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Password    string `yaml:"password"`
	Db          int    `yaml:"db"`
	RedisPrefix string `yaml:"RedisPrefix"`
}

func init() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置信息错误:{%v}", err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		log.Fatalf("配置信息解析错误:{%v}", err)
	}
	jwt := Conf.JwtSecret
	if jwt == "" {
		jwt = utils.RandString(17)
		data := map[string]interface{}{
			"JwtSecret": jwt,
		}
		for key, value := range data {
			viper.Set(key, value)
		}
		file, _ := os.OpenFile(viper.ConfigFileUsed(), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0o666)
		defer file.Close()
		viper.WriteConfigAs(file.Name())
	}

	err = viper.ReadInConfig()
	if err != nil {

	}
}
