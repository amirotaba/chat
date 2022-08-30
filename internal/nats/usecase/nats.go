package natsUsecase

import (
	"chat/domain"
	"sort"
	"strings"
)

type Usecase struct {
	NatsRepo domain.NatsRepository
	UserRepo domain.UserRepository
}

func NewUseCase(r domain.Repositories) domain.NatsUseCase {
	return &Usecase{
		NatsRepo: r.Nats,
		UserRepo: r.User,
	}
}

func (a *Usecase) JoinPv(form *domain.CreatePv) (domain.Join, error) {
	if _, err := a.UserRepo.Read(form.User2); err != nil {
		return domain.Join{}, err
	}
	var list []string
	u1 := strings.Split(form.User1, "")
	u2 := strings.Split(form.User2, "")
	list = append(list, u1[0], u2[0])
	sort.Strings(list)
	switch list[0] {
	case u1[0]:
		sub := form.User1 + "-" + form.User2
		_, err := a.NatsRepo.ReadPv(sub)
		if err != nil {
			if err = a.NatsRepo.CreatePv(domain.Private{
				Sub: sub,
			}); err != nil {
				return domain.Join{}, err
			}
		}
		pv, err := a.NatsRepo.ReadPv(sub)
		if err != nil {
			return domain.Join{}, err
		}
		msg, err := a.NatsRepo.ReadMsg(pv.ID)
		if err != nil {
			return domain.Join{}, err
		}
		return domain.Join{
			Message: msg,
			ID:      pv.ID,
		}, nil
	case u2[0]:
		sub := form.User2 + "-" + form.User1
		_, err := a.NatsRepo.ReadPv(sub)
		if err != nil {
			if err = a.NatsRepo.CreatePv(domain.Private{
				Sub: sub,
			}); err != nil {
				return domain.Join{}, err
			}
		}
		pv, err := a.NatsRepo.ReadPv(sub)
		if err != nil {
			return domain.Join{}, err
		}
		msg, err := a.NatsRepo.ReadMsg(pv.ID)
		if err != nil {
			return domain.Join{}, err
		}
		return domain.Join{
			Message: msg,
			ID:      pv.ID,
		}, nil
	}
	return domain.Join{}, nil
}

func (a *Usecase) JoinGp(name string) (domain.Join, error) {
	gp, err := a.NatsRepo.ReadGp(name)
	if err != nil {
		return domain.Join{}, err
	}
	msg, err := a.NatsRepo.ReadMsg(gp.ID)
	if err != nil {
		return domain.Join{}, err
	}
	return domain.Join{
		Message: msg,
		ID:      gp.ID,
	}, nil
}

func (a *Usecase) CreateGp(sub string) (domain.Join, error) {
	if err := a.NatsRepo.CreateGp(domain.Group{
		Sub: sub,
	}); err != nil {
		return domain.Join{}, err
	}
	gp, err := a.NatsRepo.ReadGp(sub)
	if err != nil {
		return domain.Join{}, err
	}

	return domain.Join{
		ID: gp.ID,
	}, nil
}

func (a *Usecase) CreateMsg(form domain.Message) error {
	if err := a.NatsRepo.CreateMsg(form); err != nil {
		return err
	}
	return nil
}
