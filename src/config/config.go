package config

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var envPtr = pflag.String("env", "dev", "Environment: dev or prod")

func InitLoadConfig() *AllConfig {
	// 使用pflag库来读取命令行参数，用于指定环境，默认为"dev"
	pflag.Parse()

	config := viper.New()
	// 设置读取路径
	config.AddConfigPath("./ShortService/config")
	// 设置读取文件名字
	config.SetConfigName(fmt.Sprintf("application-%s", *envPtr))
	// 设置读取文件类型
	config.SetConfigType("yaml")
	// 读取文件载体
	var configData *AllConfig
	// 读取配置文件
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Use Viper ReadInConfig Fatal error config err:%s \n", err))
	}
	// 查找对应配置文件
	err = config.Unmarshal(&configData)
	if err != nil {
		panic(fmt.Errorf("read config file to struct err: %s\n", err))
	}
	// 打印配置文件信息
	fmt.Printf("配置文件信息：%+v", configData)
	return configData
}

// AllConfig 整合Config
type AllConfig struct {
	Server     Server
	DataSource DataSource
	Log        Log
	Redis      Redis
}

type Server struct {
	Port  string
	Level string
}

type DataSource struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string `mapstructure:"db_name"`
	Config   string
}

func (d *DataSource) Dsn() string {
	return d.UserName + ":" + d.Password + "@tcp(" + d.Host + ":" + d.Port + ")/" + d.DBName + "?" + d.Config
}

type Log struct {
	Level    string
	FilePath string
}

type Redis struct {
	Host     string
	Port     string
	Password string
	DataBase int `mapstructure:"data_base"`
}
