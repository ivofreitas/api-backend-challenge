package mysql

import (
	"context"
	"database/sql"
	"github.com/sword/api-backend-challenge/api/v1/user"
	"github.com/sword/api-backend-challenge/model"
)

type mysql struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) user.Repository {
	return &mysql{db}
}

func (m *mysql) GetByUsernameAndPassword(ctx context.Context, username, password string) (*model.User, error) {
	row := m.db.QueryRowContext(ctx, `
		SELECT username, role 
		FROM sword.users 
		WHERE username = ? AND password = ?
	`, username, password)

	user := new(model.User)
	err := row.Scan(&user.Username, &user.Role)
	if err != nil {
		return nil, err
	}

	return user, nil
}
