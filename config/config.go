package config

import (
	"GuTikTok/utils"
	"github.com/spf13/viper"
	"os"
)

var Conf *Config

type Config struct {
	JwtSecret string `yaml:"jwt_secret"`
	MySQL     *Mysql `yaml:"mysql"`
	Log       *Log   `yaml:"log"`
	Redis     *Redis `yaml:"redis"`
}
type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}
type Log struct {
	Enable     bool   `yaml:"enable"`      // 是否启用日志
	Level      string `yaml:"level"`       // 日志等级，可用 panic,fatal,error,warn,info,debug,trace
	Name       string `yaml:"name"`        // 日志文件名
	MaxSize    int    `yaml:"max_size"`    // 日志最大大小
	MaxBackups int    `yaml:"max_backups"` // 日志最大备份数
	MaxAge     int    `yaml:"max_age"`     // 日志最长时间
	Compress   bool   `yaml:"compress"`    // 日志是否压缩
}
type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

func init() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {

	}
	err = viper.Unmarshal(&Conf)
	if err != nil {

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
