package user

import (
	"ocserv/internal/models"
	"ocserv/internal/repository"
	"ocserv/pkg/password"
)

func Authenticate(repository repository.UserRepositoryInterface, data CreateLoginData) (*models.User, bool) {
	user, err := repository.GetUserByUsername(data.Username)
	if err != nil {
		return nil, false
	}
	if ok := password.Compare(data.Password, user.Password); !ok {
		return nil, false
	}
	return user, false
}
