package config

import (
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Logger struct {
		Level string `yaml:"level" env:"LOG_LEVEL" env-default:"info"`
	} `yaml:"logger" env-required:"true"`
	App struct {
		Environment     string        `yaml:"environment" env:"APP_ENV" env-default:"development"`
		ShutdownTimeout time.Duration `yaml:"shutdownTimeout" env:"APP_SHUTDOWN_TIMEOUT" env-default:"5s"`
		// MyPodIP         string        `env:"MY_POD_IP"            env-default:"127.0.0.1"`
		// Domain          string        `env:"APP_DOMAIN"           env-default:"https://go-api-boilerplate.local"`
		// Secret          string        `env:"USER_SECRET"          env-default:"secret"`
		// ApiBaseURL      string        `env:"USER_BASE_URL"        env-default:"https://api.go-api-boilerplate.local/users"`
	} `yaml:"app"     env-required:"true"`
	Debug struct {
		Host string `yaml:"host" env:"DEBUG_HOST" env-default:"0.0.0.0"`
		Port int    `yaml:"port" env:"DEBUG_PORT" env-default:"4000"`
	} `yaml:"debug"   env-required:"true"`
	// DB struct {
	// 	Host     string `yaml:"host" env:"DB_HOST" env-default:"localhost" env-required:"true"`
	// 	Port     int    `yaml:"port" env:"DB_PORT" env-default:"5432" env-required:"true"`
	// 	User     string `yaml:"user" env:"DB_USER" env-required:"true"`
	// 	Password string `yaml:"password" env:"DB_PASSWORD" env-required:"true"`
	// 	Database string `yaml:"database" env:"DB_DATABASE" env-required:"true"`
	// 	SSLMode  string `yaml:"sslMode" env:"DB_SSL_MODE" env-default:"disable" env-required:"true"`
	// } `yaml:"db"    env-required:"true"`
	NATS struct {
		Host string `yaml:"host" env:"NATS_HOST" env-default:"localhost" env-required:"true"`
		Port int    `yaml:"port" env:"NATS_PORT" env-default:"4222" env-required:"true"`
	} `yaml:"nats"    env-required:"true"`
	HTTP struct {
		Host         string        `yaml:"host" env:"HTTP_HOST"`
		Port         int           `yaml:"port" env:"HTTP_PORT" env-default:"3000" env-required:"true"`
		ReadTimeout  time.Duration `yaml:"readTimeout" env:"HTTP_SERVER_READ_TIMEOUT"     env-default:"5s"`
		WriteTimeout time.Duration `yaml:"writeTimeout" env:"HTTP_SERVER_WRITE_TIMEOUT"    env-default:"10s"`
		IdleTimeout  time.Duration `yaml:"idleTimeout" env:"HTTP_SERVER_SHUTDOWN_TIMEOUT" env-default:"120s"`
	} `yaml:"http"    env-required:"true"`
	GRPC struct {
		Host          string        `yaml:"host" env:"GRPC_HOST"`
		Port          int           `yaml:"port" env:"GRPC_PORT" env-default:"3001" env-required:"true"`
		ServerMinTime time.Duration `yaml:"serverMinTime" env:"GRPC_SERVER_MIN_TIME" env-default:"5m"` // if a client pings more than once every 5 minutes (default), terminate the connection
		ServerTime    time.Duration `yaml:"serverTime" env:"GRPC_SERVER_TIME" env-default:"2h"`        // ping the client if it is idle for 2 hours (default) to ensure the connection is still active
		ServerTimeout time.Duration `yaml:"serverTimeout" env:"GRPC_SERVER_TIMEOUT" env-default:"20s"` // wait 20 second (default) for the ping ack before assuming the connection is dead
		ConnTime      time.Duration `yaml:"connTime" env:"GRPC_CONN_TIME" env-default:"10s"`           // send pings every 10 seconds if there is no activity
		ConnTimeout   time.Duration `yaml:"connTimeout" env:"GRPC_CONN_TIMEOUT" env-default:"20s"`     // wait 20 second for ping ack before considering the connection dead
	} `yaml:"grpc"    env-required:"true"`
	Tenants map[string]struct {
		Master string `yaml:"master"`
		Site   string `yaml:"site"`
		// Address string `yaml:"address"`
		DB struct {
			Host string `yaml:"host"`
			// Port     int    `yaml:"port" env-required:"true"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			Database string `yaml:"database"`
		} `yaml:"db" env-required:"true"`
	} `yaml:"tenants" env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		err := cleanenv.ReadConfig("config.yaml", instance)
		if err != nil {
			panic(err)
		}
	})
	return instance
}
