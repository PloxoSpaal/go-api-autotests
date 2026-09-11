package models

type User struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	LastName   string `json:"lastName"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
}

type UserResponse struct {
	User User `json:"user"`
}

type CreateUserRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	LastName   string `json:"lastName"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
}

type UpdateUserRequest struct {
	Email      *string `json:"email,omitempty"`
	LastName   *string `json:"lastName,omitempty"`
	FirstName  *string `json:"firstName,omitempty"`
	MiddleName *string `json:"middleName,omitempty"`
}
