package configs

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Mode   string `mapstructure:"mode"`
	Dotenv string `mapstructure:"dotenv"`
	Server struct {
		Addr            string        `mapstructure:"addr"`
		Port            string        `mapstructure:"port"`
		WriteTimeout    time.Duration `mapstructure:"write_timeout"`
		ReadTimeout     time.Duration `mapstructure:"read_timeout"`
		IdleTimeout     time.Duration `mapstructure:"idle_timeout"`
		GracefulTimeout time.Duration `mapstructure:"graceful_timeout"`
	} `mapstructure:"server"`
	CLI struct {
		Addr string `mapstructure:"addr"`
		Port string `mapstructure:"port"`
	}
}

func InitConfig() (Config, error) {
	var config Config
	v := viper.New()
	v.AddConfigPath("configs")
	v.SetConfigName("config")

	if err := v.ReadInConfig(); err != nil {
		return config, err
	}
	if err := v.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}
