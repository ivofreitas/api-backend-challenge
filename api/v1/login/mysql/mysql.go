package mysql

import (
	"database/sql"
	"github.com/sword/api-backend-challenge/api/v1/login"
	"github.com/sword/api-backend-challenge/api/v1/user"
)

type mysql struct {
	user.Repository
	db *sql.DB
}

func NewRepository(userRepository user.Repository, db *sql.DB) login.Repository {
	return &mysql{userRepository, db}
}
