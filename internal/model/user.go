package model

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type AccountStatus int

type User struct {
	ID               uuid.UUID     `validate:"required"`
	Email            string        `validate:"required,email"`
	PasswordHash     string        `validate:"required"`
	PasswordSalt     string        `validate:"required"`
	FirstName        string        `validate:"required,min=2,max=30"`
	LastName         string        `validate:"required,min=2,max=30"`
	DateOfBirth      time.Time     `validate:"required,not_future"`
	RegistrationDate time.Time     `validate:"required"`
	LastLoginDate    time.Time     `validate:"omitempty"`
	AccountStatus    AccountStatus `validate:"required"`
}

const (
	AccountStatusActive   AccountStatus = 1
	AccountStatusDisabled AccountStatus = 2
	AccountStatusDeleted  AccountStatus = 3
)

func ValidateUser(user *User) error {
	validate := validator.New()
	validate.RegisterValidation("not_future", func(fl validator.FieldLevel) bool {
		asTime, ok := fl.Field().Interface().(time.Time)
		if !ok {
			return false // it's not even a time.Time
		}
		// it's valid if the time is not after Now
		return !asTime.After(time.Now())
	})
	error := validate.Struct(user)
	if error != nil {
		return error
	}
	return nil
}

// Custom email error type
type EmailInUseError struct {
	Email string
}

func (e *EmailInUseError) Error() string {
	return fmt.Sprintf("Email %s is already in use", e.Email)
}
