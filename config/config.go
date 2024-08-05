package config

import (
	"github.com/spf13/viper"
	"os"
)

type PgConfig struct {
	Host     string
	Port     string
	Username string
	Pass     string
	DBName   string
	SSLmode  string
}

func InitConfig() error {
	// go run main.go
	//viper.AddConfigPath("../config")

	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
func PostgresConfig() *PgConfig {
	cfg := PgConfig{
		Host:     viper.GetString("test_db.host"),
		Port:     viper.GetString("test_db.port"),
		Username: viper.GetString("test_db.username"),
		Pass:     os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("test_db.database"),
		SSLmode:  viper.GetString("test_db.sslmode"),
	}
	return &cfg
}
