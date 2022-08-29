package domain

import (
	natsDomain "chat/domain/nats"
	"chat/domain/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataBase struct {
	User   *mongo.Collection
	Pv     *mongo.Collection
	Gp     *mongo.Collection
	Msg    *mongo.Collection
	Client *mongo.Client
}

type UseCases struct {
	User userDomain.UserUseCase
	Nats natsDomain.NatsUseCase
}

type Repositories struct {
	User userDomain.UserRepository
	Nats natsDomain.NatsRepository
}
