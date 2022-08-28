package domain

import (
	natsDomain "chat/domain/nats"
	"chat/domain/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataBase struct {
	Collection *mongo.Collection
	Client     *mongo.Client
}

type UseCases struct {
	User userDomain.UserUseCase
	Nats natsDomain.NatsUseCase
}

type Repositories struct {
	User userDomain.UserRepository
	Nats natsDomain.NatsRepository
}
