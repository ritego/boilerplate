package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading env file: %w", err))
	}
	viper.AutomaticEnv()
	viper.WatchConfig()
	log.Println("Config Loaded")
}

func Get(key string) interface{} {
	return viper.Get(key)
}

func String(key string) string {
	return viper.GetString(key)
}

func Integer(key string) int64 {
	return viper.GetInt64(key)
}

func Float(key string) float64 {
	return viper.GetFloat64(key)
}

func Boolean(key string) bool {
	return viper.GetBool(key)
}

func Duration(key string) time.Duration {
	return viper.GetDuration(key)
}
