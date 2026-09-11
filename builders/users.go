package builders

import (
	coursev1 "github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/PloxoSpaal/go-api-autotests/models"
)

type UserCreate struct {
	Email      string
	Password   string
	LastName   string
	FirstName  string
	MiddleName string
}

type UserCreateOption func(*UserCreate)

func (b *Builder) UserCreate(options ...UserCreateOption) UserCreate {
	user := UserCreate{
		Email:      b.generator.EmailWithDomain("example.com"),
		Password:   b.generator.Password(),
		LastName:   b.generator.LastName(),
		FirstName:  b.generator.FirstName(),
		MiddleName: b.generator.MiddleName(),
	}

	for _, option := range options {
		option(&user)
	}

	return user
}

func WithUserCreateEmail(email string) UserCreateOption {
	return func(user *UserCreate) { user.Email = email }
}

func WithUserCreatePassword(password string) UserCreateOption {
	return func(user *UserCreate) { user.Password = password }
}

func WithUserCreateLastName(lastName string) UserCreateOption {
	return func(user *UserCreate) { user.LastName = lastName }
}

func WithUserCreateFirstName(firstName string) UserCreateOption {
	return func(user *UserCreate) { user.FirstName = firstName }
}

func WithUserCreateMiddleName(middleName string) UserCreateOption {
	return func(user *UserCreate) { user.MiddleName = middleName }
}

func (u UserCreate) HTTPRequest() models.CreateUserRequest {
	return models.CreateUserRequest{
		Email:      u.Email,
		Password:   u.Password,
		LastName:   u.LastName,
		FirstName:  u.FirstName,
		MiddleName: u.MiddleName,
	}
}

func (u UserCreate) GRPCRequest() *coursev1.CreateUserRequest {
	return &coursev1.CreateUserRequest{
		Email:      u.Email,
		Password:   u.Password,
		LastName:   u.LastName,
		FirstName:  u.FirstName,
		MiddleName: u.MiddleName,
	}
}

type UserUpdate struct {
	Email      *string
	LastName   *string
	FirstName  *string
	MiddleName *string
}

type UserUpdateOption func(*UserUpdate)

func (b *Builder) UserUpdate(options ...UserUpdateOption) UserUpdate {
	email := b.generator.EmailWithDomain("example.com")
	lastName := b.generator.LastName()
	firstName := b.generator.FirstName()
	middleName := b.generator.MiddleName()

	update := UserUpdate{
		Email:      &email,
		LastName:   &lastName,
		FirstName:  &firstName,
		MiddleName: &middleName,
	}

	for _, option := range options {
		option(&update)
	}

	return update
}

func WithUserUpdateEmail(email string) UserUpdateOption {
	return func(update *UserUpdate) { update.Email = &email }
}

func WithUserUpdateLastName(lastName string) UserUpdateOption {
	return func(update *UserUpdate) { update.LastName = &lastName }
}

func WithUserUpdateFirstName(firstName string) UserUpdateOption {
	return func(update *UserUpdate) { update.FirstName = &firstName }
}

func WithUserUpdateMiddleName(middleName string) UserUpdateOption {
	return func(update *UserUpdate) { update.MiddleName = &middleName }
}

func (u UserUpdate) HTTPRequest() models.UpdateUserRequest {
	return models.UpdateUserRequest{
		Email:      u.Email,
		LastName:   u.LastName,
		FirstName:  u.FirstName,
		MiddleName: u.MiddleName,
	}
}

func (u UserUpdate) GRPCRequest(userID string) *coursev1.UpdateUserRequest {
	return &coursev1.UpdateUserRequest{
		Id:         userID,
		Email:      u.Email,
		LastName:   u.LastName,
		FirstName:  u.FirstName,
		MiddleName: u.MiddleName,
	}
}
