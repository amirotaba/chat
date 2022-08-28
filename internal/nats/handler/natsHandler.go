package natsHandler

import (
	"chat/domain/nats"
	"chat/utils"
	"fmt"
)

type Handler struct {
	UseCase natsDomain.NatsUseCase
}

func NewHandler(u natsDomain.NatsUseCase) *Handler {
	return &Handler{
		UseCase: u,
	}
}

func (m *Handler) JoinPv() string {
	fmt.Println("enter the username of user that you want to send message: ")
	username := utils.Read()
	message, err := m.UseCase.JoinPv(username)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(message)
	return username
}

func (m *Handler) JoinGp() string {
	fmt.Println("enter the group name of group that you want to join: ")
	name := utils.Read()

	message, err := m.UseCase.JoinGp(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(message)
	return name
}

func (m *Handler) CreateGp() string {
	fmt.Println("enter the group name that you want to create: ")
	name := utils.Read()

	m.UseCase.CreateGp(name)
	return name
}
