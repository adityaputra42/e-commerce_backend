package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DbDriver             string        `mapstructure:"DB_DRIVER"`
	DbSource             string        `mapstructure:"DB_SOURCE"`
	ServerAddress        string        `mapstructure:"ADDRESS_SERVER"`
	SecretKey            string        `mapstructure:"SECRET_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATIOn"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
