package config

import (
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	ApiKey string
	Port   string
}

func NewEnv() *Env {
	viper.SetDefault("API_KEY", "SeCrEt824")
	viper.SetDefault("PORT", ":8080")
	env := Env{}

	if env.ApiKey = os.Getenv("API_KEY"); env.ApiKey == "" {
		env.ApiKey = viper.GetString("API_KEY")
	}
	if env.Port = os.Getenv("PORT"); env.Port == "" {
		env.Port = viper.GetString("PORT")
	}
	return &env
}
