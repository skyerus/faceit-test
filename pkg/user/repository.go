package user

import (
	"github.com/skyerus/faceit-test/pkg/customerror"
)

// Repository - handles data transfer between the application and the database
type Repository interface {
	Create(u *User) customerror.Error
	Get(ID int) (User, customerror.Error)
}
