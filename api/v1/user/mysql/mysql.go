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

func (m *mysql) Create(ctx context.Context, user *model.User) error {

	insert := `
	INSERT INTO sword.users(id, username, password, role_id, created_at, updated_at) 
	VALUES (?, ?, ?, (SELECT id FROM sword.roles WHERE position = ?), NOW(), NOW())`

	_, err := m.db.ExecContext(
		ctx,
		insert,
		user.ID,
		user.Username,
		user.Password,
		user.Role)

	return err
}

func (m *mysql) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	row := m.db.QueryRowContext(ctx, `
		SELECT u.username, u.password, r.position 
		FROM sword.users u
		JOIN sword.roles r ON u.role_id = r.id 
		WHERE u.username = ?
	`, username)

	user := new(model.User)
	err := row.Scan(&user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	return user, nil
}
