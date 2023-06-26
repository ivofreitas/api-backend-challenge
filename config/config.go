package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
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
	Broker string `mapstructure:"broker"`
	Topic  string `mapstructure:"topic"`
}

var (
	env  *Env
	once sync.Once
)

// GetEnv returns env values
func GetEnv() *Env {

	once.Do(func() {

		viper.AutomaticEnv()
		if err := godotenv.Load("config/.env"); err != nil {
			log.Println(err)
		}

		env = new(Env)
		env.Server.Host = viper.GetString("SERVER_HOST")
		env.Server.BasePath = viper.GetString("SERVER_BASE_PATH")
		env.Server.Port = viper.GetString("SERVER_PORT")

		env.Authorization.Secret = viper.GetString("AUTHORIZATION_SECRET")

		env.Log.Enabled = viper.GetBool("LOG_ENABLED")
		env.Log.Level = viper.GetString("LOG_LEVEL")

		env.Doc.Title = viper.GetString("DOC_TITLE")
		env.Doc.Description = viper.GetString("DOC_DESCRIPTION")
		env.Doc.Enabled = viper.GetBool("DOC_ENABLED")
		env.Doc.Version = viper.GetString("DOC_VERSION")

		env.MySQL.Username = viper.GetString("MYSQL_USERNAME")
		env.MySQL.Password = viper.GetString("MYSQL_PASSWORD")
		env.MySQL.Host = viper.GetString("MYSQL_HOST")
		env.MySQL.Database = viper.GetString("MYSQL_DATABASE")
		env.MySQL.PoolConn = viper.GetInt("MYSQL_POOL_CONN")
		env.MySQL.ConnLifetime = viper.GetDuration("MYSQL_CONN_LIFETIME")

		env.Kafka.Broker = viper.GetString("KAFKA_BROKER")
		env.Kafka.Topic = viper.GetString("KAFKA_TOPIC")

	})
	return env
}
