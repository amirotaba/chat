package userUseCase

import (
	"chat/domain"
	"errors"
)

type Usecase struct {
	UserRepo domain.UserRepository
}

func NewUseCase(r domain.Repositories) domain.UserUseCase {
	return &Usecase{
		UserRepo: r.User,
	}
}

func (a *Usecase) Create(user domain.User) error {
	res, _ := a.UserRepo.Read(user.UserName)
	if res.UserName != "" {
		return errors.New("this username is taken")
	}

	if err := a.UserRepo.Create(user); err != nil {
		return err
	}

	return nil
}

func (a *Usecase) SignIn(user domain.User) error {
	u, err := a.UserRepo.Read(user.UserName)
	if err != nil {
		return err
	}

	if user.PassWord != u.PassWord {
		return errors.New("wrong password")
	}

	return nil
}
