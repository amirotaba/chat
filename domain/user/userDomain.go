package userDomain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserName string             `json:"user_name"`
	PassWord string             `json:"pass_word"`
}

type UserRepository interface {
	Create(form User) error
	Read(username string) (User, error)
}

type UserUseCase interface {
	Create(user User) error
	SignIn(user User) error
}

type UserHandler interface {
	SignIn() (User, error)
	SignUp() error
}
