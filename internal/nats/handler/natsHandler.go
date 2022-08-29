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

func (m *Handler) JoinPv(username string) natsDomain.Join {
	fmt.Println("enter the username of user that you want to send message: ")
	user2 := utils.Read()
	form, err := m.UseCase.JoinPv(&natsDomain.CreatePv{
		User1: username,
		User2: user2,
	})
	if err != nil {
		fmt.Println(err)
	}
	form.Sub = user2
	for i := range form.Message {
		fmt.Println(form.Message[i].Message)
	}
	return form
}

func (m *Handler) JoinGp() natsDomain.Join {
	fmt.Println("enter the group name of group that you want to join: ")
	name := utils.Read()

	form, err := m.UseCase.JoinGp(name)
	if err != nil {
		fmt.Println(err)
	}
	form.Sub = name
	for i := range form.Message {
		fmt.Println(form.Message[i].Message)
	}
	return form
}

func (m *Handler) CreateGp() (natsDomain.Join, error) {
	fmt.Println("enter the group name that you want to create: ")
	name := utils.Read()

	msg, err := m.UseCase.CreateGp(name)
	if err != nil {
		return natsDomain.Join{}, err
	}
	msg.Sub = name
	return msg, nil
}

func (m *Handler) CreateMsg(form natsDomain.Message) error {
	if err := m.UseCase.CreateMsg(form); err != nil {
		return err
	}
	return nil
}
