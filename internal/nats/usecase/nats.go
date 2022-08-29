package natsUsecase

import (
	"chat/domain"
	natsDomain "chat/domain/nats"
	"sort"
	"strings"
)

type Usecase struct {
	NatsRepo natsDomain.NatsRepository
}

func NewUseCase(r domain.Repositories) natsDomain.NatsUseCase {
	return &Usecase{
		NatsRepo: r.Nats,
	}
}

func (a *Usecase) JoinPv(form *natsDomain.CreatePv) (natsDomain.Join, error) {
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
			if err = a.NatsRepo.CreatePv(natsDomain.Private{
				Sub: sub,
			}); err != nil {
				return natsDomain.Join{}, err
			}
		}
		pv, err := a.NatsRepo.ReadPv(sub)
		if err != nil {
			return natsDomain.Join{}, err
		}
		msg, err := a.NatsRepo.ReadMsg(pv.ID)
		if err != nil {
			return natsDomain.Join{}, err
		}
		return natsDomain.Join{
			Message: msg,
			ID:      pv.ID,
		}, nil
	case u2[0]:
		sub := form.User2 + "-" + form.User1
		_, err := a.NatsRepo.ReadPv(sub)
		if err != nil {
			if err = a.NatsRepo.CreatePv(natsDomain.Private{
				Sub: sub,
			}); err != nil {
				return natsDomain.Join{}, err
			}
		}
		pv, err := a.NatsRepo.ReadPv(sub)
		if err != nil {
			return natsDomain.Join{}, err
		}
		msg, err := a.NatsRepo.ReadMsg(pv.ID)
		if err != nil {
			return natsDomain.Join{}, err
		}
		return natsDomain.Join{
			Message: msg,
			ID:      pv.ID,
		}, nil
	}
	return natsDomain.Join{}, nil
}

func (a *Usecase) JoinGp(name string) (natsDomain.Join, error) {
	gp, err := a.NatsRepo.ReadGp(name)
	if err != nil {
		return natsDomain.Join{}, err
	}
	msg, err := a.NatsRepo.ReadMsg(gp.ID)
	if err != nil {
		return natsDomain.Join{}, err
	}
	return natsDomain.Join{
		Message: msg,
		ID:      gp.ID,
	}, nil
}

func (a *Usecase) CreateGp(sub string) (natsDomain.Join, error) {
	if err := a.NatsRepo.CreateGp(natsDomain.Group{
		Sub: sub,
	}); err != nil {
		return natsDomain.Join{}, err
	}
	gp, err := a.NatsRepo.ReadGp(sub)
	if err != nil {
		return natsDomain.Join{}, err
	}

	return natsDomain.Join{
		ID: gp.ID,
	}, nil
}

func (a *Usecase) CreateMsg(form natsDomain.Message) error {
	if err := a.NatsRepo.CreateMsg(form); err != nil {
		return err
	}
	return nil
}
