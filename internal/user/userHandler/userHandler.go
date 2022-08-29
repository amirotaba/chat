package userHandler

import (
	"chat/domain/user"
	"chat/utils"
	"fmt"
)

type Handler struct {
	UseCase userDomain.UserUseCase
}

func NewHandler(u userDomain.UserUseCase) userDomain.UserHandler {
	return &Handler{
		UseCase: u,
	}
}

func (m *Handler) SignIn() (userDomain.User, error) {
	var user userDomain.User
	fmt.Println("SignIn: ")
	fmt.Println("enter your username: ")
	user.UserName = utils.Read()
	fmt.Println("enter your password: ")
	user.PassWord = utils.Read()

	if err := m.UseCase.SignIn(user); err != nil {
		return userDomain.User{}, err
	}
	return user, nil
}

func (m *Handler) SignUp() error {
	var user userDomain.User
	fmt.Println("enter your name: ")
	user.UserName = utils.Read()
	fmt.Println("enter your password: ")
	user.PassWord = utils.Read()

	if err := m.UseCase.Create(user); err != nil {
		return err
	}
	return nil
}
