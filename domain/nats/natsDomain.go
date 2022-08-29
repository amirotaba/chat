package natsDomain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Message struct {
	ChatID  primitive.ObjectID
	Sub     string
	Message string
	Time    time.Time
}

type Group struct {
	ID  primitive.ObjectID `bson:"_id,omitempty"`
	Sub string
}

type Private struct {
	ID  primitive.ObjectID `bson:"_id,omitempty"`
	Sub string
}

type CreatePv struct {
	User1 string
	User2 string
}

type Join struct {
	Message []Message
	ID      primitive.ObjectID
	Sub     string
}

type NatsRepository interface {
	CreatePv(form Private) error
	ReadPv(sub string) (Private, error)
	CreateGp(form Group) error
	ReadGp(sub string) (Group, error)
	CreateMsg(form Message) error
	ReadMsg(id primitive.ObjectID) ([]Message, error)
}

type NatsUseCase interface {
	CreateGp(sub string) (Join, error)
	JoinPv(form *CreatePv) (Join, error)
	JoinGp(sub string) (Join, error)
	CreateMsg(form Message) error
}
