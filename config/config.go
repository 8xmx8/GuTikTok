package main

import (
	"GuTikTok/utils"
	"github.com/spf13/viper"
	"os"
)

var Conf *Config

type Config struct {
	JwtSecret string `yaml:"jwt_secret"`
	MySQL     *Mysql `yaml:"mysql"`
}
type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

func main() {
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
