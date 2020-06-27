package user

import (
	"github.com/skyerus/faceit-test/pkg/customerror"
)

// Service - high level API for performing actions on user
type Service interface {
	Create(u *User) customerror.Error
	Get(ID int) (User, customerror.Error)
	Delete(ID int) customerror.Error
	GetAll(f Filter) ([]User, customerror.Error)
	Update(u User) customerror.Error
}
