package userservice

import (
	"github.com/skyerus/faceit-test/pkg/customerror"
	"github.com/skyerus/faceit-test/pkg/user"
)

type userService struct {
	userRepo user.Repository
}

// NewUserService ...
func NewUserService(userRepo user.Repository) user.Service {
	return &userService{userRepo}
}

func (us userService) Create(u *user.User) customerror.Error {
	return us.userRepo.Create(u)
}

func (us userService) Get(ID int) (user.User, customerror.Error) {
	return us.userRepo.Get(ID)
}

func (us userService) Delete(ID int) customerror.Error {
	return us.userRepo.Delete(ID)
}
