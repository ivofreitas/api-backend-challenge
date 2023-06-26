package config

import (
	"github.com/spf13/viper"
	"sync"
	"time"
)

// Env values
type Env struct {
	Server        Server        `mapstructure:"server"`
	Authorization Authorization `mapstructure:"authorization"`
	Log           Log           `mapstructure:"log"`
	Doc           Doc           `mapstructure:"doc"`
	MySQL         MySQL         `mapstructure:"mysql"`
	Kafka         Kafka         `mapstructure:"kafka"`
}

// Server config
type Server struct {
	Host     string `mapstructure:"host"`
	BasePath string `mapstructure:"base_path"`
	Port     string `mapstructure:"port"`
}

type Authorization struct {
	Secret string `mapstructure:"secret"`
}

// Log config
type Log struct {
	Enabled bool   `mapstructure:"enabled"`
	Level   string `mapstructure:"level"`
}

// Doc - swagger information
type Doc struct {
	Title       string `mapstructure:"title"`
	Description string `mapstructure:"description"`
	Enabled     bool   `mapstructure:"enabled"`
	Version     string `mapstructure:"version"`
}

// MySQL - db conn information
type MySQL struct {
	Username     string        `mapstructure:"username"`
	Password     string        `mapstructure:"password"`
	Host         string        `mapstructure:"host"`
	Database     string        `mapstructure:"database"`
	PoolConn     int           `mapstructure:"pool_conn"`
	ConnLifetime time.Duration `mapstructure:"conn_lifetime"`
}

// Kafka - message broker information
type Kafka struct {
	Brokers []string `mapstructure:"brokers"`
	Topic   string   `mapstructure:"topic"`
}

var (
	env  *Env
	once sync.Once
)

// GetEnv returns env values
func GetEnv() *Env {

	once.Do(func() {
		viper.AddConfigPath("./config")
		viper.SetConfigName("config")
		viper.SetConfigType("json")

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			return
		}

		viper.Unmarshal(&env)
	})
	return env
}
