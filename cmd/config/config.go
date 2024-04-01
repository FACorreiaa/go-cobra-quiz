package config

import "github.com/spf13/viper"

func InitConfig() error {
	v := viper.New()
	v.AddConfigPath("config")
	v.AddConfigPath(".")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.WriteConfig(); err != nil {
		return err
	}
	return nil

}
