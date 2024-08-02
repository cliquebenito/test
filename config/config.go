package config

import "github.com/spf13/viper"

func InitConfig() error {
	// если  go run main.go
	//viper.AddConfigPath("../config")

	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
