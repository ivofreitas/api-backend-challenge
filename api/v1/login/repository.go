package login

import "github.com/sword/api-backend-challenge/api/v1/user"

type Repository interface {
	user.Repository
}
