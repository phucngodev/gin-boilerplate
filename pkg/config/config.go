package config

import (
	"fmt"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

var ProviderSet = wire.NewSet(New)

type Config struct {
	Mode         string      `mapstructure:"mode"`
	Port         string      `mapstructure:"port"`
	AppName      string      `mapstructure:"appName"`
	Url          string      `mapstructure:"url"`
	MaxPingCount int         `mapstructure:"maxPingCount"`
	DBConfig     DBConfig    `mapstructure:"database"`
	RedisConfig  RedisConfig `mapstructure:"redis"`
	LogConfig    LogConfig   `mapstructure:"log"`
}

type DBConfig struct {
	Dbname          string `mapstructure:"dbname"`
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	MaximumPoolSize int    `mapstructure:"maximumPoolSize"`
	MaximumIdleSize int    `mapstructure:"maximumIdleSize"`
	LogMode         bool   `mapstructure:"logMode"`
}

type RedisConfig struct {
	Addr         string `mapstructure:"address"`
	Password     string `mapstructure:"password"`
	Db           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"poolSize"`
	MinIdleConns int    `mapstructure:"minIdleConns"`
	IdleTimeout  int    `mapstructure:"idleTimeout"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	TimeFormat string `mapstructure:"timeFormat"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
	Compress   bool   `mapstructure:"compress"`
	LocalTime  bool   `mapstructure:"localTime"`
	Stdout     bool   `mapstructure:"stdout"`
}

func New(path string) (*Config, error) {
	viper.SetConfigType("yaml")
	if path != "" {
		viper.SetConfigFile(path)
	} else {
		viper.AddConfigPath("config")
		viper.SetConfigFile("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("error umarshal config: %w", err)
	}

	return config, nil
}
