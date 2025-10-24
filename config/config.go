package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppName  string `mapstructure:"APP_NAME"`
	Port     string `mapstructure:"PORT"`
	Env      string `mapstructure:"ENV"`
	LogLevel string `mapstructure:"LOG_LEVEL"`
}

func Load() *Config {
	viper.SetConfigName("config") // config.yaml or config.json
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	// Defaults
	viper.SetDefault("APP_NAME", "go-microservice")
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("ENV", "development")
	viper.SetDefault("LOG_LEVEL", "info")

	// Environment overrides
	viper.AutomaticEnv()

	// Try reading file (ignore if missing)
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No config file found, using defaults and environment variables")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Config unmarshal error: %v", err)
	}

	return &cfg
}
