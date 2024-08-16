package types

type UserStore interface {
    GetUserByEmail(email string) (*User, error)
    GetUserByID(id string) (*User, error)
    CreateUser(User) error
}

type User struct {
    ID       string
    Name     string
    Email    string
    Password string
    CreatedAt string
    UpdatedAt string
}

type LoginUserPayload struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}
