package config

import (
	"path/filepath"
	"time"
	"runtime"

	"github.com/spf13/viper"
)

type Env struct {
	ContextTimeout     time.Duration
	ApiKey   string `mapstructure:"API_KEY"`
	Port     string `mapstructure:"PORT"`
	RedisURL string `mapstructure:"REDIS_URL"`
}

func NewEnv() *Env {
	viper.SetDefault("API_KEY", "SeCrEt824")
	viper.SetDefault("PORT", ":8084")
	viper.SetDefault("CONTEXT_TIMEOUT", 30)
	viper.SetDefault("REDIS_URL", "redis://localhost:6379/0?")

	_, file, _, ok := runtime.Caller(1)
	if ok {
		viper.SetConfigFile(filepath.Join(filepath.Dir(file), ".env"))
	}
	_ = viper.ReadInConfig() // Ignore error if config file not found
	viper.AutomaticEnv()

	var cfg Env
	if err := viper.Unmarshal(&cfg); err != nil {
		panic("Failed to unmarshal config: " + err.Error())
	}
	cfg.ContextTimeout = time.Duration(viper.GetInt64("CONTEXT_TIMEOUT") * int64(time.Second))

	return &cfg
}
