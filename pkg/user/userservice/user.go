package userservice

import (
	"github.com/skyerus/faceit-test/pkg/crypto"
	"github.com/skyerus/faceit-test/pkg/customerror"
	"github.com/skyerus/faceit-test/pkg/event"
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
	saltedHash, err := crypto.GenerateSaltedHash(u.Password)
	if err != nil {
		return customerror.NewGenericHTTPError(err)
	}
	u.Password = saltedHash
	customErr := us.userRepo.Create(u)
	if customErr != nil {
		return customErr
	}
	go event.BroadcastCreateEvent(*u)
	return nil
}

func (us userService) Get(ID int) (user.User, customerror.Error) {
	return us.userRepo.Get(ID)
}

func (us userService) Delete(ID int) customerror.Error {
	customErr := us.userRepo.Delete(ID)
	if customErr != nil {
		return customErr
	}
	go event.BroadcastDeleteEvent(ID)
	return nil
}

func (us userService) GetAll(f user.Filter) ([]user.User, customerror.Error) {
	return us.userRepo.GetAll(f)
}

func (us userService) Update(u user.User) customerror.Error {
	saltedHash, err := crypto.GenerateSaltedHash(u.Password)
	if err != nil {
		return customerror.NewGenericHTTPError(err)
	}
	u.Password = saltedHash
	customErr := us.userRepo.Update(u)
	if customErr != nil {
		return customErr
	}
	go event.BroadcastUpdateEvent(u)
	return nil
}
