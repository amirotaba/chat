package natsUsecase

import (
	"chat/domain"
	natsDomain "chat/domain/nats"
)

type Usecase struct {
	NatsRepo natsDomain.NatsRepository
}

func NewUseCase(r domain.Repositories) natsDomain.NatsUseCase {
	return &Usecase{
		NatsRepo: r.Nats,
	}
}

func (a *Usecase) JoinPv(username string) ([]natsDomain.Message, error) {
	m, err := a.NatsRepo.Read(username)
	if err != nil {
		if err = a.NatsRepo.Create(natsDomain.Message{
			Sub: username,
		}); err != nil {
			return nil, err
		}
	}
	return m, nil
}

func (a *Usecase) JoinGp(name string) ([]natsDomain.Message, error) {
	m, err := a.NatsRepo.Read(name)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (a *Usecase) CreateGp(sub string) error {
	if err := a.NatsRepo.Create(natsDomain.Message{
		Sub: sub,
	}); err != nil {
		return err
	}
	return nil
}
