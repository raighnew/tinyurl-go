package setting

import (
	"fmt"
	"log"
	"time"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Server struct {
	Environment  string        `env:"Environment" envDefault:"test"`
	RunMode      string        `env:"RUN_MODE" envDefault:"debug"`
	HttpPort     int           `env:"PORT" envDefault:"8081"`
	BaseUrl      string        `env:"BASE_URL" envDefault:"http://localhost:8081/"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" envDefault:"10s"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" envDefault:"10s"`
}

var ServerSetting = &Server{}

type SQLDatabase struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var SQLDatabaseSetting = &SQLDatabase{}

type Redis struct {
	Host        string        `env:"REDIS_HOST" envDefault:"localhost:6379"`
	Password    string        `env:"REDIS_PASSWORD"`
	MaxIdle     int           `env:"REDIS_MAX_IDLE" envDefault:"9000"`
	MaxActive   int           `env:"REDIS_MAX_ACTIVE" envDefault:"10000"`
	IdleTimeout time.Duration `env:"REDIS_IDLE_TIMEOUT"`
}

var RedisSetting = &Redis{}

func Setup() {
	err := godotenv.Load() // 👈 load .env file
	if err != nil {
		log.Fatal(err)
	}
	loadEnv(ServerSetting)
	loadEnv(SQLDatabaseSetting)
	loadEnv(RedisSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// mapTo map section
func loadEnv(cfg interface{}) {
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
}
