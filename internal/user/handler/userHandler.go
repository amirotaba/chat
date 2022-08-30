package userHandler

import (
	"chat/domain"
	"chat/utils"
	"fmt"
)

type Handler struct {
	UseCase domain.UserUseCase
}

func NewHandler(u domain.UserUseCase) domain.UserHandler {
	return &Handler{
		UseCase: u,
	}
}

func (m *Handler) SignIn() (domain.User, error) {
	var user domain.User
	fmt.Println("SignIn: ")
	fmt.Println("enter your username: ")
	user.UserName = utils.Read()
	fmt.Println("enter your password: ")
	user.PassWord = utils.Read()

	if err := m.UseCase.SignIn(user); err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (m *Handler) SignUp() error {
	var user domain.User
	fmt.Println("enter your name: ")
	user.UserName = utils.Read()
	fmt.Println("enter your password: ")
	user.PassWord = utils.Read()

	if err := m.UseCase.Create(user); err != nil {
		return err
	}
	return nil
}
