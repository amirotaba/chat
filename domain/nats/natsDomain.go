package natsDomain

import "time"

type Message struct {
	Sub     string
	Message string
	Time    time.Time
}

type NatsRepository interface {
	Create(message Message) error
	Read(sub string) ([]Message, error)
}

type NatsUseCase interface {
	CreateGp(sub string) error
	JoinPv(sub string) ([]Message, error)
	JoinGp(sub string) ([]Message, error)
}

type NatsHandler interface {
	JoinPv()
	JoinGp()
	Create()
}
