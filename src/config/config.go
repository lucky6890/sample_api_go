package config

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Port    string
	RunMode string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SslMode  bool
}

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	Db                 int8
	MinIdleConnections int
	PoolSize           int
	PoolTimeout        int
}

func GetConfig() *Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := loadConfig(cfgPath, "yml")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := parseConfig(v)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

func parseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
		return nil, err
	}
	return &cfg, nil
}

func loadConfig(filename string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Unable to load config: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file Not Found")
		}
		return nil, err
	}
	return v, nil
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "../../src/config/config-docker.yml"
	} else if env == "production" {
		return "../../src/config/config-production.yml"
	} else {
		return "../../src/config/config-development.yml"
	}
}
