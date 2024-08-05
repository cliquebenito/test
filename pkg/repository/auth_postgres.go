package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"test/internal/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var u models.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE email = $1 AND password = $2", models.UserTable)
	err := r.db.Get(&u, query, username, password)
	return u, err

}
func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email,password) values ($1,$2) RETURNING id", models.UserTable)
	row := r.db.QueryRow(query, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
