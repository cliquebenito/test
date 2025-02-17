package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"test/config"
)

func NewPostgres(cfg *config.PgConfig) (*sqlx.DB, error) {

	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Pass, cfg.SSLmode))
	if err != nil {
		return nil, err
	}

	return db, nil

}
